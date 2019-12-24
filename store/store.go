// Package store handlings all the storage in GCSE backend.
package store

import (
	"log"
	"time"

	"github.com/daviddengcn/bolthelper"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golangplus/bytes"
	"github.com/golangplus/errors"

	"github.com/x0rzkov/gcse/configs"
	gpb "github.com/x0rzkov/gcse/shared/proto"
)

var (
	// pkgs
	//   - <site>
	//     - <path> -> PackageInfo
	// persons
	//   - <site>
	//     - <id> -> PersonInfo
	// history
	//  - pkgs
	//     - <site/path> -> HistoryInfo
	//  - persons
	//     - <site/id> -> HistoryInfo
	// repos
	//  - <site>
	//    - <user>
	//     - <repo> -> Repository
	pkgsRoot    = []byte("pkgs")
	personsRoot = []byte("persons")
	historyRoot = []byte("history")
	reposRoot   = []byte("repos")
)

var box = &bh.RefCountBox{
	DataPath: configs.StoreBoltPath,
}

func RepoInfoAge(r *gpb.RepoInfo) time.Duration {
	t, _ := ptypes.Timestamp(r.CrawlingTime)
	return time.Now().Sub(t)
}

// Returns all the sites one by one by calling the provided func.
func ForEachPackageSite(f func(string) error) error {
	return box.View(func(tx bh.Tx) error {
		return tx.ForEach([][]byte{pkgsRoot}, func(_ bh.Bucket, k, v bytesp.Slice) error {
			if v != nil {
				log.Printf("Unexpected value %q for key %q, ignored", string(v), string(k))
				return nil
			}
			return errorsp.WithStacks(f(string(k)))
		})
	})
}

func ForEachPackageOfSite(site string, f func(string, *gpb.PackageInfo) error) error {
	return box.View(func(tx bh.Tx) error {
		return tx.ForEach([][]byte{pkgsRoot, []byte(site)}, func(_ bh.Bucket, k, v bytesp.Slice) error {
			if v == nil {
				log.Printf("Unexpected nil value for key %q, ignored", string(k))
				return nil
			}
			info := &gpb.PackageInfo{}
			if err := errorsp.WithStacksAndMessage(proto.Unmarshal(v, info), "Unmarshal %d bytes failed", len(v)); err != nil {
				log.Printf("Unmarshal failed: %v, ignored", err)
				return nil
			}
			return errorsp.WithStacks(f(string(k), info))
		})
	})
}

// Returns an empty (non-nil) PackageInfo if not found.
func ReadPackage(site, path string) (*gpb.PackageInfo, error) {
	info := &gpb.PackageInfo{}
	if err := box.View(func(tx bh.Tx) error {
		return tx.Value([][]byte{pkgsRoot, []byte(site), []byte(path)}, func(bs bytesp.Slice) error {
			if err := errorsp.WithStacksAndMessage(proto.Unmarshal(bs, info), "Unmarshal %d bytes failed", len(bs)); err != nil {
				log.Printf("Unmarshal failed: %v", err)
				*info = gpb.PackageInfo{}
			}
			return nil
		})
	}); err != nil {
		return nil, err
	}
	return info, nil
}

func UpdatePackage(site, path string, f func(*gpb.PackageInfo) error) error {
	return box.Update(func(tx bh.Tx) error {
		b, err := tx.CreateBucketIfNotExists([][]byte{pkgsRoot, []byte(site)})
		if err != nil {
			return err
		}
		info := &gpb.PackageInfo{}
		if err := b.Value([][]byte{[]byte(path)}, func(bs bytesp.Slice) error {
			if err := errorsp.WithStacksAndMessage(proto.Unmarshal(bs, info), "Unmarshal %d bytes", len(bs)); err != nil {
				log.Printf("Unmarshaling failed: %v", err)
				*info = gpb.PackageInfo{}
			}
			return nil
		}); err != nil {
			return err
		}
		if err := errorsp.WithStacks(f(info)); err != nil {
			return err
		}
		bs, err := proto.Marshal(info)
		if err != nil {
			return errorsp.WithStacksAndMessage(err, "marshaling %v failed: %v", info, err)
		}
		return b.Put([][]byte{[]byte(path)}, bs)
	})
}

func DeletePackage(site, path string) error {
	return box.Update(func(tx bh.Tx) error {
		return tx.Delete([][]byte{pkgsRoot, []byte(site), []byte(path)})
	})
}

func ReadPerson(site, id string) (*gpb.PersonInfo, error) {
	info := &gpb.PersonInfo{}
	if err := box.View(func(tx bh.Tx) error {
		return tx.Value([][]byte{personsRoot, []byte(site), []byte(id)}, func(bs bytesp.Slice) error {
			if err := errorsp.WithStacksAndMessage(proto.Unmarshal(bs, info), "Unmarshal %d bytes failed", len(bs)); err != nil {
				log.Printf("Unmarshal failed: %v", err)
				*info = gpb.PersonInfo{}
			}
			return nil
		})
	}); err != nil {
		return nil, err
	}
	return info, nil
}

func UpdatePerson(site, id string, f func(*gpb.PersonInfo) error) error {
	return box.Update(func(tx bh.Tx) error {
		b, err := tx.CreateBucketIfNotExists([][]byte{personsRoot, []byte(site)})
		if err != nil {
			return err
		}
		info := &gpb.PersonInfo{}
		if err := b.Value([][]byte{[]byte(id)}, func(bs bytesp.Slice) error {
			err := errorsp.WithStacksAndMessage(proto.Unmarshal(bs, info), "Unmarshal %d bytes", len(bs))
			if err != nil {
				log.Printf("Unmarshaling failed: %v", err)
				*info = gpb.PersonInfo{}
			}
			return nil
		}); err != nil {
			return err
		}
		if err := errorsp.WithStacks(f(info)); err != nil {
			return err
		}
		bs, err := proto.Marshal(info)
		if err != nil {
			return errorsp.WithStacksAndMessage(err, "marshaling %v failed: %v", info, err)
		}
		return b.Put([][]byte{[]byte(id)}, bs)
	})
}

func DeletePerson(site, id string) error {
	return box.Update(func(tx bh.Tx) error {
		return tx.Delete([][]byte{personsRoot, []byte(site), []byte(id)})
	})
}

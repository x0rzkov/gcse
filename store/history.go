package store

import (
	"log"
	"time"

	"github.com/daviddengcn/bolthelper"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golangplus/bytes"
	"github.com/golangplus/errors"

	gpb "github.com/x0rzkov/gcse/shared/proto"
)

func SaveSnapshot(path string) error {
	return box.Update(func(tx bh.Tx) error {
		return tx.CopyFile(path, 0644)
	})
}

const (
	maxHistoryEvents = 10
)

func readHistoryOf(box *bh.RefCountBox, root []byte, site, idOrPath string) (*gpb.HistoryInfo, error) {
	info := &gpb.HistoryInfo{}
	if err := box.View(func(tx bh.Tx) error {
		return tx.Value([][]byte{historyRoot, root, []byte(site), []byte(idOrPath)}, func(bs bytesp.Slice) error {
			if err := errorsp.WithStacksAndMessage(proto.Unmarshal(bs, info), "Unmarshal %d bytes failed", len(bs)); err != nil {
				log.Printf("Unmarshal failed: %v", err)
				*info = gpb.HistoryInfo{}
			}
			return nil
		})
	}); err != nil {
		return nil, err
	}
	return info, nil
}

func readHistory(root []byte, site, idOrPath string) (*gpb.HistoryInfo, error) {
	return readHistoryOf(box, root, site, idOrPath)
}

func ReadPackageHistory(site, path string) (*gpb.HistoryInfo, error) {
	return readHistory(pkgsRoot, site, path)
}

func ReadPackageHistoryOf(box *bh.RefCountBox, site, path string) (*gpb.HistoryInfo, error) {
	return readHistoryOf(box, pkgsRoot, site, path)
}

func ReadPersonHistory(site, path string) (*gpb.HistoryInfo, error) {
	return readHistory(personsRoot, site, path)
}

func updateHistory(root []byte, site, idOrPath string, f func(*gpb.HistoryInfo) error) error {
	return box.Update(func(tx bh.Tx) error {
		b, err := tx.CreateBucketIfNotExists([][]byte{historyRoot, root, []byte(site)})
		if err != nil {
			return err
		}
		info := &gpb.HistoryInfo{}
		if err := b.Value([][]byte{[]byte(idOrPath)}, func(bs bytesp.Slice) error {
			err := errorsp.WithStacksAndMessage(proto.Unmarshal(bs, info), "Unmarshal %d bytes", len(bs))
			if err != nil {
				log.Printf("Unmarshaling failed: %v", err)
				*info = gpb.HistoryInfo{}
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
		return b.Put([][]byte{[]byte(idOrPath)}, bs)
	})
}

func UpdatePackageHistory(site, path string, f func(*gpb.HistoryInfo) error) error {
	return updateHistory(pkgsRoot, site, path, f)
}

func AppendPackageEvent(site, path, foundWay string, t time.Time, a gpb.HistoryEvent_Action_Enum) error {
	return UpdatePackageHistory(site, path, func(hi *gpb.HistoryInfo) error {
		if hi.FoundTime == nil {
			// The first time the package was found
			hi.FoundTime, _ = ptypes.TimestampProto(t)
			hi.FoundWay = foundWay
		}
		if a == gpb.HistoryEvent_Action_None {
			return nil
		}
		// Insert the event
		tsp, _ := ptypes.TimestampProto(t)
		hi.Events = append([]*gpb.HistoryEvent{{
			Action:    a,
			Timestamp: tsp,
		}}, hi.Events...)
		if len(hi.Events) > maxHistoryEvents {
			hi.Events = hi.Events[:maxHistoryEvents]
		}
		switch a {
		case gpb.HistoryEvent_Action_Success:
			hi.LatestSuccess = tsp
		case gpb.HistoryEvent_Action_Failed:
			hi.LatestFailed = tsp
		}
		return nil
	})
}

func UpdatePersonHistory(site, path string, f func(*gpb.HistoryInfo) error) error {
	return updateHistory(personsRoot, site, path, f)
}

func deleteHistory(root []byte, site, idOrPath string) error {
	return box.Update(func(tx bh.Tx) error {
		return tx.Delete([][]byte{historyRoot, root, []byte(site), []byte(idOrPath)})
	})
}

func DeletePackageHistory(site, path string) error {
	return deleteHistory(pkgsRoot, site, path)
}

func DeletePersonHistory(site, path string) error {
	return deleteHistory(personsRoot, site, path)
}

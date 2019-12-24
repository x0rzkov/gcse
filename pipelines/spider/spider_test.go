package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golangplus/testing/assert"
	"github.com/golangplus/time"

	"github.com/x0rzkov/gcse/configs"
	gpb "github.com/x0rzkov/gcse/shared/proto"
	"github.com/x0rzkov/gcse/spider/github"
	"github.com/x0rzkov/gcse/store"
)

func init() {
	configs.SetTestingDataPath()
}

func TestShouldCrawlLater(t *testing.T) {
	assert.True(t, "res", shouldCrawlLater(&RepositoryInfo{
		Repository: &gpb.Repository{
			CrawlingInfo: &gpb.CrawlingInfo{},
		},
	}, &RepositoryInfo{
		Repository: &gpb.Repository{},
	}))
	assert.False(t, "res", shouldCrawlLater(&RepositoryInfo{
		Repository: &gpb.Repository{},
	}, &RepositoryInfo{
		Repository: &gpb.Repository{
			CrawlingInfo: &gpb.CrawlingInfo{},
		},
	}))
}

func TestSelectRepos(t *testing.T) {
	const (
		site  = "TestSelectRepos.com"
		user1 = "david"
		name1 = "hello"
		user2 = "daviddeng"
		name2 = "world"
		user3 = "daviddeng"
		name3 = "go"
	)
	assert.NoError(t, store.UpdateRepository(site, user1, name1, func(r *gpb.Repository) error {
		r.Stars = 1
		return nil
	}))
	now := time.Now()
	assert.NoError(t, store.UpdateRepository(site, user2, name2, func(r *gpb.Repository) error {
		r.Stars = 2
		r.CrawlingInfo = &gpb.CrawlingInfo{}
		r.CrawlingInfo.SetCrawlingTime(now.Add(-10 * timep.Day))
		return nil
	}))
	ts3, _ := ptypes.TimestampProto(now.Add(-15 * timep.Day))
	assert.NoError(t, store.UpdateRepository(site, user3, name3, func(r *gpb.Repository) error {
		r.Stars = 3
		r.CrawlingInfo = &gpb.CrawlingInfo{
			CrawlingTime: ts3,
		}
		return nil
	}))

	repos, err := selectRepos(site, 2)
	assert.NoError(t, err)
	assert.Equal(t, "repos", repos, []*RepositoryInfo{{
		Repository: &gpb.Repository{
			Stars: 1,
		},
		Name: name1,
		User: user1,
	}, {
		Repository: &gpb.Repository{
			Stars: 3,
			CrawlingInfo: &gpb.CrawlingInfo{
				CrawlingTime: ts3,
			},
		},
		Name: name3,
		User: user3,
	}})
}

func TestCrawlRepo_UnknownSite(t *testing.T) {
	ctx := context.Background()
	assert.Error(t, crawlRepo(ctx, "unknown.com", nil))
}

func initSpider() {
	githubSpider = github.NewSpiderWithContents(map[string]string{
		"/repos/daviddengcn/gcse/branches/master": `{
			"commit": {
				"sha": "sha-1"
			}
		}`, "/repos/daviddengcn/gcse/git/trees/sha-1?recursive=1": `{
			"tree": [
		      {
			    "path": "a.go",
			    "type": "blob",
			    "sha": "sha-2"
			  }
			],
			"truncated": false
		}`, "/repos/daviddengcn/gcse/contents/a.go": `{
			"name": "bi.go",
			"path": "bi.go",
			"sha": "sha-2",
			"content": "cGFja2FnZSBnY3NlCgppbXBvcnQgKAoJImdpdGh1Yi5jb20vZGF2aWRkZW5n\nY24vZ28tZWFzeWJpIgopCgpmdW5jIEFkZEJpVmFsdWVBbmRQcm9jZXNzKGFn\nZ3IgYmkuQWdncmVnYXRlTWV0aG9kLCBuYW1lIHN0cmluZywgdmFsdWUgaW50\nKSB7CgliaS5BZGRWYWx1ZShhZ2dyLCBuYW1lLCB2YWx1ZSkKCWJpLkZsdXNo\nKCkKCWJpLlByb2Nlc3MoKQp9Cg==\n",
			"encoding": "base64",
			"type": "file"
		}`, "/repos/daviddengcn/unchanged/branches/master": `{
			"commit": {
				"sha": "sha-unchanged"
			}
		}`})
}

func TestCrawlRepo(t *testing.T) {
	ctx := context.Background()
	tm := time.Now()
	now = timep.PresetNow(tm)
	const (
		user = "daviddengcn"
		repo = "gcse"
	)
	initSpider()
	r := &RepositoryInfo{
		Repository: &gpb.Repository{
			Branch:       "master",
			CrawlingInfo: (&gpb.CrawlingInfo{}).SetCrawlingTime(tm.Add(-timep.Day)),
		},
		User: user,
		Name: repo,
	}
	assert.NoError(t, crawlRepo(ctx, "github.com", r))
	assert.Equal(t, "r.Repository", *r.Repository, gpb.Repository{
		Branch:    "master",
		Signature: "sha-1",
		Packages: map[string]*gpb.Package{
			"": {
				Name:        "gcse",
				Path:        "",
				Imports:     []string{"github.com/daviddengcn/go-easybi"},
				TestImports: nil,
			}},
		CrawlingInfo: (&gpb.CrawlingInfo{}).SetCrawlingTime(tm),
	})
}

func TestCrawlRepo_Unchanged(t *testing.T) {
	ctx := context.Background()
	tm := time.Now()
	now = timep.PresetNow(tm)
	initSpider()
	r := &RepositoryInfo{
		Repository: &gpb.Repository{
			Branch:    "master",
			Signature: "sha-unchanged",
			Packages: map[string]*gpb.Package{
				"": {
					Name:        "gcse",
					Path:        "",
					Imports:     []string{"github.com/daviddengcn/go-easybi"},
					TestImports: nil,
				}},
		},
		User: "daviddengcn",
		Name: "unchanged",
	}
	assert.NoError(t, crawlRepo(ctx, "github.com", r))
	assert.Equal(t, "r.Repository", *r.Repository, gpb.Repository{
		Branch:    "master",
		Signature: "sha-unchanged",
		Packages: map[string]*gpb.Package{
			"": {
				Name:        "gcse",
				Path:        "",
				Imports:     []string{"github.com/daviddengcn/go-easybi"},
				TestImports: nil,
			}},
		CrawlingInfo: (&gpb.CrawlingInfo{}).SetCrawlingTime(tm),
	})
}

func TestCrawlAndSaveRepo_RepositoryDeleted(t *testing.T) {
	ctx := context.Background()
	tm := time.Now()
	now = timep.PresetNow(tm)
	const (
		site = "github.com"
		user = "daviddengcn"
		repo = "nonexist"
	)
	initSpider()
	r := &RepositoryInfo{
		Repository: &gpb.Repository{
			Branch:       "master",
			CrawlingInfo: (&gpb.CrawlingInfo{}).SetCrawlingTime(tm.Add(-timep.Day)),
		},
		User: user,
		Name: repo,
	}
	assert.NoError(t, store.UpdateRepository(site, user, repo, func(doc *gpb.Repository) error {
		*doc = *r.Repository
		return nil
	}))
	assert.NoError(t, crawlAndSaveRepo(ctx, site, r))

	rp, err := store.ReadRepository(site, user, repo)
	assert.NoError(t, err)
	assert.Equal(t, "rp", rp, &gpb.Repository{})
}

func cleanDatabase(t *testing.T) {
	assert.NoErrorOrDie(t, os.RemoveAll(configs.StoreBoltPath()))
}

func TestExec(t *testing.T) {
	tm := time.Now()
	now = timep.PresetNow(tm)
	const (
		site = "github.com"
		user = "daviddengcn"
		repo = "gcse"
	)
	cleanDatabase(t)
	assert.NoError(t, store.UpdateRepository(site, user, repo, func(r *gpb.Repository) error {
		r.Branch = "master"
		return nil
	}))
	initSpider()
	assert.NoError(t, exec(1, time.Hour))
	r, err := store.ReadRepository(site, user, repo)
	assert.NoError(t, err)
	assert.Equal(t, "r", *r, gpb.Repository{
		Branch:    "master",
		Signature: "sha-1",
		Packages: map[string]*gpb.Package{
			"": {
				Name:        "gcse",
				Path:        "",
				Imports:     []string{"github.com/daviddengcn/go-easybi"},
				TestImports: nil,
			}},
		CrawlingInfo: (&gpb.CrawlingInfo{}).SetCrawlingTime(tm),
	})
}

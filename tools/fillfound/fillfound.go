package main

import (
	"log"

	"github.com/golangplus/time"

	"github.com/x0rzkov/gcse"
	"github.com/x0rzkov/gcse/shared/proto"
	"github.com/x0rzkov/gcse/store"
	"github.com/x0rzkov/gcse/utils"
)

func doFill() error {
	cDB := gcse.LoadCrawlerDB()
	return cDB.PackageDB.Iterate(func(pkg string, val interface{}) error {
		ent, ok := val.(gcse.CrawlingEntry)
		if !ok {
			log.Printf("Wrong entry, ignored: %+v", ent)
			return nil
		}
		site, path := utils.SplitPackage(pkg)
		return store.AppendPackageEvent(site, path, "unknown", ent.ScheduleTime.Add(-10*timep.Day), gcsepb.HistoryEvent_Action_None)
	})
}

func main() {
	if err := doFill(); err != nil {
		log.Fatalf("doFill failed: %v", err)
	}
}

package main

import (
	"context"
	"flag"
	"net"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	"github.com/x0rzkov/gcse/configs"
	gpb "github.com/x0rzkov/gcse/shared/proto"
	"github.com/x0rzkov/gcse/store"
	"github.com/x0rzkov/gcse/utils"
)

type server struct {
}

var _ gpb.StoreServiceServer = (*server)(nil)

func (s *server) PackageCrawlHistory(_ context.Context, req *gpb.PackageCrawlHistoryReq) (*gpb.PackageCrawlHistoryResp, error) {
	site, path := utils.SplitPackage(req.Package)
	info, err := store.ReadPackageHistory(site, path)
	if err != nil {
		glog.Errorf("ReadPackageHistoryOf %q %q failed: %v", site, path, err)
		return nil, err
	}
	return &gpb.PackageCrawlHistoryResp{Info: info}, nil
}

func main() {
	addr := flag.String("addr", configs.StoreDAddr, "addr to listen")

	flag.Parse()

	glog.Infof("Listening to %s", *addr)
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	gpb.RegisterStoreServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}

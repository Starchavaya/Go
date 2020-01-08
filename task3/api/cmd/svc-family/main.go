package main

import (
	"github.com/Starchavaya/Go/task3/api/internal/conn"
	"github.com/Starchavaya/Go/task3/api/internal/swc"
	"github.com/Starchavaya/Go/task3/api/middleware"
	pb "github.com/Starchavaya/Go/task3/api/proto"
	mw "github.com/grpc-ecosystem/go-grpc-middleware" // ...
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//protoc --go_out=plugins=grpc:. service.proto
	//protoc -I$GOPATH/bin -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. service.proto
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(mw.WithUnaryServerChain(middleware.LoggerMiddleware)) //
	pb.RegisterIMotherServer(s, swc.NewFamilyServer(conn.DbConnection{}.GetConnection()))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

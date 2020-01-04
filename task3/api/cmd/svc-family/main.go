package main

import (
	"github.com/Starchavaya/Go/task3/api/internal/conn"
	"github.com/Starchavaya/Go/task3/api/internal/swc"
	pb "github.com/Starchavaya/Go/task3/api/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"github.com/Starchavaya/Go/task3/api/middleware"
    mw "github.com/grpc-ecosystem/go-grpc-middleware" // ...
)

func main() {
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

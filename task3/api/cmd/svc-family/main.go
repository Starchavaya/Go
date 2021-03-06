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
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Println("failed to listen: %v", err)
	}
	s := grpc.NewServer(mw.WithUnaryServerChain(middleware.LoggerMiddleware))
	pb.RegisterIMotherServer(s, swc.NewFamilyServer(conn.DbConnection{}.GetConnection()))
	if err := s.Serve(lis); err != nil {
		log.Println("failed to serve: %v", err)
	}
}

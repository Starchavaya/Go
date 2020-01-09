package main

import (
	"context"
	lib "github.com/Starchavaya/Go/task3/api/proto"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"
)

func Run(address string, opts ...runtime.ServeMuxOption) error {
	gw := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	//err := lib.RegisterIMotherHandlerFromEndpoint(context.Background(), gw, "localhost:9091", dialOpts)
	err := lib.RegisterIMotherHandlerFromEndpoint(context.Background(), gw, "svc:9091", dialOpts)
	if err != nil {
		return err;
	}

	mux := http.NewServeMux()
	mux.Handle("/", gw)
	return http.ListenAndServe(address, gw)
}

func main() {
	defer glog.Flush()
	//8080
	if err := Run("gw:8080"); err != nil {
		glog.Fatal(err)
	}
}

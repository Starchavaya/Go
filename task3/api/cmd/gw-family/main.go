package main

import (
	"context"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"
	gw "github.com/Starchavaya/Go/task3/api/proto"
)

func Run(address string,opts ...runtime.ServeMuxOption)error  {
	mux:=runtime.NewServeMux()
	mux2:=http.NewServeMux()
	dialOpts:=[]grpc.DialOption{grpc.WithInsecure()}
	err:=gw.RegisterIMotherHandlerFromEndpoint(context.Background(),mux,"localhost:9091",dialOpts)
	if err!=nil{
		return err;
	}
	mux2.Handle("/",mux)
	return http.ListenAndServe(address,mux)
}

func main()  {
	defer glog.Flush()
	if err:=Run(":8080");err!=nil{
		glog.Fatal(err)
	}
}

package main

import (
	"fmt"
	pb "golang_test/golang_grpc/protobuf/grpc_learning"
	pb_iml "golang_test/golang_grpc/protobuf_implement"
	"net"

	"google.golang.org/grpc"
)

func main() {

	ipaddr := "127.0.0.1:8877"
	listen, err := net.Listen("tcp", ipaddr)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, pb_iml.NewGreeterObj())
	fmt.Printf("Start server...")
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to start serve: %v", err)
	}
}

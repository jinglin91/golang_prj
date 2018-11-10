package main

import (
	"fmt"

	pb "golang_test/golang_grpc/protobuf/grpc_learning"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	addr = "127.0.0.1:8877"
	name = "Gogrpc"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("failed to dial: %v\n", err)
		return
	}
	fmt.Printf("connected to server(%s)\n", addr)
	clientObj := pb.NewGreeterClient(conn)
	rsp, err := clientObj.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		fmt.Printf("failed to call Sayhello: %v", err)
		return
	}
	fmt.Printf("server rsp: %s\n", rsp.Message)
}

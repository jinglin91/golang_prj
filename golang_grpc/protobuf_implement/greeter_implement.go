package protobuf_implement

import (
	"fmt"
	pb "golang_test/golang_grpc/protobuf/grpc_learning"

	"golang.org/x/net/context"
)

func NewGreeterObj() (s *greeterImplement) {
	greeterObj := new(greeterImplement)
	return greeterObj
}

type greeterImplement struct{}

func (s *greeterImplement) SayHello(ct context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("#####get client request name:", in.Name, "#####")
	return &pb.HelloResponse{Message: "hello " + in.Name}, nil
}

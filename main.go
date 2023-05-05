package main

import (
	"context"
	"log"
	"net"

	helloworldv1 "github.com/pranayhere/grpc-gw-2/rpc/hw/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	helloworldv1.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldv1.HelloRequest) (*helloworldv1.HelloReply, error) {
	return &helloworldv1.HelloReply{Message: "hello " + in.Name, Timestamp: timestamppb.Now()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: ", err)
	}

	s := grpc.NewServer()
	helloworldv1.RegisterGreeterServer(s, NewServer())

	log.Fatal(s.Serve(lis))
}

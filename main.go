package main

import (
	"context"
	"log"
	"net"

	helloworldv2 "github.com/pranayhere/grpc-gw-2/proto/proto/helloworld"
	"google.golang.org/grpc"
)

type server struct {
	helloworldv2.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldv2.HelloRequest) (*helloworldv2.HelloReply, error) {
	return &helloworldv2.HelloReply{Message: "hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: ", err)
	}

	s := grpc.NewServer()
	helloworldv2.RegisterGreeterServer(s, NewServer())

	log.Fatal(s.Serve(lis))
}

package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterGreetServiceServer(srv, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	log.Println("server hello")
	time.Sleep(31 * time.Minute)
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}

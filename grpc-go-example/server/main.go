package main

import (
	"context"
	"log"
	"net"

	"grpc-go-example/proto"

	"google.golang.org/grpc"
)

// server implements the GreeterServer interface
type server struct {
	proto.UnimplementedGreeterServer
}

// SayHello handles the SayHello RPC method
func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	// Set up a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the Greeter service with the gRPC server
	proto.RegisterGreeterServer(s, &server{})

	// Start serving requests
	log.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

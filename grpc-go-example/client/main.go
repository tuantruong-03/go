package main

import (
	"context"
	"log"
	"time"

	"grpc-go-example/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Create a context with a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Establish a connection to the gRPC server
	conn, err := grpc.DialContext(
		ctx,
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()), // Use insecure credentials for non-TLS
		grpc.WithBlock(), // Block until the connection is established
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new Greeter client
	c := proto.NewGreeterClient(conn)

	// Call the SayHello method
	resp, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: " World!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Print the response message
	log.Printf("Greeting: %s", resp.GetMessage())
}

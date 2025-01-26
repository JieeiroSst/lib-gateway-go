package main

import (
	"context"
	"log"
	"time"

	pb "github.com/JIeeiroSst/lib-gateway-go/gateway/proto"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewLibServiceClient(conn)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make RPC call
	response, err := client.GetUser(ctx, &pb.GetUserRequest{
		UserId: "123",
	})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	// Process response
	log.Printf("User: %s, Email: %s", response.Name, response.Email)
}

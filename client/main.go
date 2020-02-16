package main

import(
	"context"
	"log"
	"time"
	"google.golang.org/grpc"
	pb "proto"
)

const (
	address     = "localhost:4040"
)

func main(){

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAddServiceClient(conn)
	
	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.Add(ctx, &pb.Request{A: 5, B:2})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetResult())

	r, err = c.Multiply(ctx, &pb.Request{A: 5, B:6})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetResult())
}
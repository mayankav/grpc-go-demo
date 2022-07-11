package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/mayankav/grpc-go-demo/greet/proto"
)

var addr string = "localhost:5051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", addr)
	}

	log.Printf("Connected to: %v\n", addr)

	defer conn.Close()
	//..
	client := pb.NewGreetServiceClient(conn)

	// from ./greet.go
	doGreet(client)
}

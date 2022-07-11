package main

import (
	"log"

	pb "github.com/mayankav/grpc-go-demo/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:5052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", addr)
	}

	log.Printf("Connected to: %v\n", addr)

	defer conn.Close()
	//..
	client := pb.NewCalcServiceClient(conn)

	// from ./greet.go
	findSum(client)
}

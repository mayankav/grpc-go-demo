package main

import (
	"context"
	"log"

	pb "github.com/mayankav/grpc-go-demo/greet/proto"
)

func doGreet(client pb.GreetServiceClient) {
	log.Println("doGreet was invoked..")
	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Mayank",
	})

	if err != nil {
		log.Fatalf("Could not greet! %v\n", err)
	}

	log.Printf("Greeting! %s\n", res.Result)
}

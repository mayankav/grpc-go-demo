package main

import (
	"context"
	"log"

	pb "github.com/mayankav/grpc-go-demo/calculator/proto"
)

func findSum(client pb.CalcServiceClient) {
	log.Println("findSum was invoked...")
	res, err := client.Sum(context.Background(), &pb.CalcRequest{
		Num1: 3,
		Num2: 10,
	})

	if err != nil {
		log.Fatalf("Could not find sum! %v\n", err)
	}

	log.Printf("The sum is %v\n", res.Result)
}

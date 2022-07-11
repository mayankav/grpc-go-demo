package main

import (
	"context"
	"io"
	"log"

	pb "github.com/mayankav/grpc-go-demo/greet/proto"
)

func doGreetManyTimes(client pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked..")

	req := &pb.GreetRequest{
		FirstName: "Mayank",
	}
	// calling the RPC endpoint
	stream, err := client.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not greet many times! %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not greet many times. Error while reading the stream! %v\n", err)
		}

		log.Printf("Greeting many times! %s\n", msg.Result)
	}

}

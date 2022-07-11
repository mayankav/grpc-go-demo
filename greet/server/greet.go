package main

import (
	"context"
	"log"

	pb "github.com/mayankav/grpc-go-demo/greet/proto"
)

// implementation of Greet RPC Api Endpoints from greet/proto/greet_grpc.pb.go
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResposne, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResposne{
		Result: "Hello " + in.FirstName,
	}, nil
}

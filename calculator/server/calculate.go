package main

import (
	"context"
	"log"

	pb "github.com/mayankav/grpc-go-demo/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.CalcRequest) (*pb.CalcResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.CalcResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}

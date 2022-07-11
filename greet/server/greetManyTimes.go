package main

import (
	"fmt"
	"log"

	pb "github.com/mayankav/grpc-go-demo/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes was invoked with: %v\n", in)
	for i := 0; i < 15; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)
		stream.Send(&pb.GreetResposne{
			Result: res,
		})
	}

	// error
	return nil
}

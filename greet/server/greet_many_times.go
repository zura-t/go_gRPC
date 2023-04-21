package main

import (
	"fmt"
	"log"

	pb "github.com/zura-t/go_gRPC/greet/proto"
)

func (server *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Println(in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
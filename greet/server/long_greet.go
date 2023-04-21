package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/zura-t/go_gRPC/greet/proto"
)

func (server *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	res := ""
	for {
		req, err := stream.Recv()
		
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error: ", err)
		}
		
		res += fmt.Sprintf("Hello %s! ", req.FirstName)
	}
}
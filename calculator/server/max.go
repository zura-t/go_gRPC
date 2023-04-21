package main

import (
	"io"
	"log"

	pb "github.com/zura-t/go_gRPC/calculator/proto"
)

func (server *Server) Max(stream pb.CalculatorService_MaxServer) error {
	var maximum int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		if number := req.Number; number > maximum {
			maximum = number
			err := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})
			if err != nil {
				log.Fatalf("Error: %v\n", err)
			}
		}
	}
}
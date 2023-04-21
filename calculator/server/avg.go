package main

import (
	"io"
	"log"

	pb "github.com/zura-t/go_gRPC/calculator/proto"
)

func (server *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	var sum int32 = 0
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("Error: ", err)
		}
		sum += req.Number
		count++
	}
}
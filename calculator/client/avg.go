package main

import (
	"context"
	"log"

	pb "github.com/zura-t/go_gRPC/calculator/proto"
)

func doAvg(x pb.CalculatorServiceClient) {
	stream, err := x.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error: ", err)
	}
	numbers := []int32{3, 5, 9, 54, 23}
	for _, number := range numbers {
		log.Println(number)
		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	log.Println(res)
}
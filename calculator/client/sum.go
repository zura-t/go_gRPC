package main

import (
	"context"
	"log"

	pb "github.com/zura-t/go_gRPC/calculator/proto"
)

func sum(x pb.CalculatorServiceClient) {
	res, err := x.Sum(context.Background(), &pb.SumRequest{
		Number1: 23,
		Number2: 23,
	})
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	log.Printf("%d\n", res.Result)
}
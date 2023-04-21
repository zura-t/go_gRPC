package main

import (
	"context"
	"io"
	"log"

	pb "github.com/zura-t/go_gRPC/calculator/proto"
)

func divideToPrimes(x pb.CalculatorServiceClient) {
	req := &pb.PrimesRequest{
		Number: 1234567890,
	}
	stream, err := x.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		log.Printf("%d\n", res.Result)
	}
}
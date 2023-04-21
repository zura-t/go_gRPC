package main

import (
	"context"
	"log"

	pb "github.com/zura-t/go_gRPC/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(x pb.CalculatorServiceClient, n int32) {
	res, err := x.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %v\n", e.Message())
			log.Printf("Error code from server: %v\n", e.Code())
			if e.Code() == codes.InvalidArgument {
				log.Println("You sent a negative number")
				return
			}
		} else {
			log.Fatalf("Non gRPC error: ", err)
		}
	}
	log.Printf("%f\n", res.Result)

}
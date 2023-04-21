package main

import (
	"context"
	"log"
	"time"

	pb "github.com/zura-t/go_gRPC/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(x pb.GreetServiceClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Ryan",
	}
	res, err := x.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
				return
			}
		} else {
			log.Fatalf("Non gRPC error: ", err)
		}
	}
	log.Printf(res.Result)
}

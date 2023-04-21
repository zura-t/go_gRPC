package main

import (
	"context"
	"log"
	"time"

	pb "github.com/zura-t/go_gRPC/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println(in)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, status.Error(codes.Canceled, "Time is finished")
		}
		time.Sleep(1 * time.Second)
	}
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

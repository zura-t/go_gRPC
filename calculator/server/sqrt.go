package main

import (
	"context"
	"fmt"
	"math"

	pb "github.com/zura-t/go_gRPC/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	number := in.Number
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", number),
		)
	}
	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
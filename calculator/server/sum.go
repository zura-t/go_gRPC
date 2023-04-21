package main

import (
	"context"

	pb "github.com/zura-t/go_gRPC/calculator/proto"
)

func (server Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Result: in.Number1 + in.Number2,
	}, nil
}



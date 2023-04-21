package main

import (
	pb "github.com/zura-t/go_gRPC/calculator/proto"
)

func (server *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	number := in.Number
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimesResponse{
				Result: divisor,
			})
			number /= divisor
		} else {
			divisor++
		}
	}
	return nil
}
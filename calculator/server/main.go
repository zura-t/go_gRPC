package main

import (
	"log"
	"net"

	pb "github.com/zura-t/go_gRPC/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:4567"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	server := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(server, &Server{})
	reflection.Register(server)

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
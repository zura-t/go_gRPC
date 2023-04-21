package main

import (
	"log"
	"net"

	pb "github.com/zura-t/go_gRPC/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:3456"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	options := []grpc.ServerOption{}
	tls := false

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		sslCreds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatal(err)
		}
		options = append(options, grpc.Creds(sslCreds))
	}

	server := grpc.NewServer(options...)
	pb.RegisterGreetServiceServer(server, &Server{})

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

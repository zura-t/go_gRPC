package main

import (
	"log"
	// "time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/zura-t/go_gRPC/greet/proto"
)

var addr string = "localhost:3456"

func main() {
	tls := false
	options := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatal(err)
		}

		options = append(options, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, options...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	x := pb.NewGreetServiceClient(conn)

	doGreet(x)
	// doGreetManyTimes(x)
	// doLongGreet(x)
	// doGreetEveryone(x)
	// doGreetWithDeadline(x, 5*time.Second)
	// doGreetWithDeadline(x, 1*time.Second)
}

package main

import (
	"context"
	"log"
	"time"

	pb "github.com/zura-t/go_gRPC/greet/proto"
)

func doLongGreet(x pb.GreetServiceClient) {
	reqs := []*pb.GreetRequest {
		{FirstName: "Ryan Gosling"},
		{FirstName: "Test"},
	}
	stream, err := x.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error: ", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error: ", err)
	}

	log.Printf(res.Result)
}
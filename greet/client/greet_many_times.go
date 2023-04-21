package main

import (
	"context"
	"io"
	"log"

	pb "github.com/zura-t/go_gRPC/greet/proto"
)

func doGreetManyTimes(x pb.GreetServiceClient) {
	req := &pb.GreetRequest{
		FirstName: "Zura",
	}

	stream, err := x.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error in stream: %v ", err)
		}

		log.Printf("%s\n", msg.Result)
	}
}

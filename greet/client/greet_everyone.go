package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/zura-t/go_gRPC/greet/proto"
)

func doGreetEveryone(x pb.GreetServiceClient) {
	stream, err := x.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	reqs := []*pb.GreetRequest{
		{FirstName: "Test"},
		{FirstName: "Second Test"},
	}
	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error: %v\n", err)
				break
			}
			log.Println(res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
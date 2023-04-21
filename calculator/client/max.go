package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/zura-t/go_gRPC/calculator/proto"
)

func doMax(x pb.CalculatorServiceClient) {
	stream, err := x.Max(context.Background())
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	waitc := make(chan struct{})
	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
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
				log.Fatalf("Error: %v\n", err)
				break
			}
			log.Println(res)
			log.Println(res.Result)
		}
		close(waitc)
	}()
	<-waitc
}

package main

import (
	"context"
	"log"

	pb "github.com/zura-t/go_gRPC/greet/proto"
)

func doGreet(x pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := x.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Zura",
	})

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Printf("%s\n", res.Result)
}
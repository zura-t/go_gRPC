package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/zura-t/go_gRPC/calculator/proto"
)

var addr string = "localhost:4567"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	x := pb.NewCalculatorServiceClient(conn)

	// sum(x)
	// divideToPrimes(x)
	// doAvg(x)
	// doMax(x)
	doSqrt(x, 25)
	doSqrt(x, -22)
}

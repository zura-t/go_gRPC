package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/zura-t/go_gRPC/blog/proto"
)

var addr string = "localhost:4567"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)

	id := createBlog(client)
	readBlog(client, id)
	// readBlog(client, "123")
	// updateBlog(client, id)
	// listBlog(client)
	deleteBlog(client, id)
}

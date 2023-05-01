package main

import (
	"context"
	"log"

	pb "github.com/zura-t/go_gRPC/blog/proto"
)

func deleteBlog(client pb.BlogServiceClient, id string) {
	_, err := client.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Blog was deleted")
}
package main

import (
	"context"
	"log"

	pb "github.com/zura-t/go_gRPC/blog/proto"
)

func updateBlog(client pb.BlogServiceClient, id string) {
	newBlog := &pb.Blog{
		Id: id,
		AuthorId: "Test",
		Title: "A new title",
		Content: "Content of the first blog with some additions",
	}

	_, err := client.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Blog was updated!")
}
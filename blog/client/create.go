package main

import (
	"context"
	"log"

	pb "github.com/zura-t/go_gRPC/blog/proto"
)

func createBlog(client pb.BlogServiceClient) string {
	blog := &pb.Blog{
		AuthorId: "Zura",
		Title: "My First blog",
		Content: "Content of the first blog",
	}

	res, err := client.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Id)
	return res.Id
}
package main

import (
	"context"
	"log"

	pb "github.com/zura-t/go_gRPC/blog/proto"
)

func readBlog(client pb.BlogServiceClient, id string) *pb.Blog {
	req := &pb.BlogId{Id: id}
	res, err := client.ReadBlog(context.Background(), req)

	if err != nil {
		log.Println(err)
	}

	log.Print(res)
	return res
}
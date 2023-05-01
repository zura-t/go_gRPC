package main

import (
	"context"
	"io"
	"log"

	pb "github.com/zura-t/go_gRPC/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(client pb.BlogServiceClient) {
	stream, err := client.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatal(err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Println(res)
	}
}
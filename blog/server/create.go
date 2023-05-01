package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/zura-t/go_gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Print(in)

	data := BlogItem {
		AuthorId: in.AuthorId,
		Title: in.Title,
		Content: in.Content,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprint(err),
		)
	}

	objectID, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Can't convert to ObjectID",
		)
	}

	return &pb.BlogId{
		Id: objectID.Hex(),
	}, nil
}
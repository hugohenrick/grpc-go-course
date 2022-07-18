package main

import (
	"context"
	"io"
	"log"

	pb "github.com/hugohenrick/grpc-go-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("---listBlog was invoked---")

	stream, err := c.ListBlogId(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Printf("Error happened while ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something Happened: %v\n", err)
		}

		log.Println(res)
	}

}

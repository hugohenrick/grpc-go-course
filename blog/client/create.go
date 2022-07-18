package main

import (
	"context"
	"log"

	pb "github.com/hugohenrick/grpc-go-course/blog/proto"
)

func CreateBlog(c pb.BlogServiceClient) string {
	log.Println("---CreateBlog was invoked---")

	blog := &pb.Blog{
		AuthorId: "Hugo Henrick",
		Title:    "My First Blog",
		Content:  "Content of the my first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}

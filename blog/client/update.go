package main

import (
	"context"
	"log"

	pb "github.com/hugohenrick/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Hugo Henrick",
		Title:    "A new title",
		Content:  "Content of the first blog updated",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Printf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}

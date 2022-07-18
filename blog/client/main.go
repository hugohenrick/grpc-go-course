package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/hugohenrick/grpc-go-course/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)

	id := CreateBlog(c)
	readBlog(c, id) //Valid
	//readBlog(c, "aNomExistingId") //Invalid
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)

}

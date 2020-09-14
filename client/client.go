package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	postpb "grpcBlog/pb"
	"io"
)

func main() {
	fmt.Println("Starting blog client")

	dial := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", dial)
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	client := postpb.NewPostServiceClient(cc)
	rsp, err := client.CreatePost(context.Background(), &postpb.CreatePostRequest{Post: &postpb.Post{
		AuthorId: "1",
		Title:    "Test",
		Content:  "Testing awesome mongo with GRPC",
	}})
	if err != nil {
		panic(err)
	}
	fmt.Printf("post created %v", rsp)

	stream, err := client.ListPost(context.Background(), &postpb.ListPostRequest{})
	if err != nil {
		fmt.Println("List error!")
		panic(err)
	}
	for {
		postRsp, readErr := stream.Recv()
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			panic(readErr)
		}
		fmt.Println(postRsp.GetPost())
	}
}

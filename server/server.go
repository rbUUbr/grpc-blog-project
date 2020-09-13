package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	blogpb "grpcBlog/pb"
	"net"
	"os"
	"os/signal"
	"time"
)

var collection *mongo.Collection

type server struct{}
type post struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func main() {
	fmt.Println("started blog server")

	// connect to mongo
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection = client.Database("blog").Collection("posts")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})
	go func() {
		fmt.Println("starting server")
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("exit... ")
	s.Stop()
	client.Disconnect(ctx)
	lis.Close()
}

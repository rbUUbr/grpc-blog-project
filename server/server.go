package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	postpb "grpcBlog/pb"
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
	postpb.RegisterPostServiceServer(s, &server{})
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

func (s *server) CreatePost(ctx context.Context, request *postpb.CreatePostRequest) (*postpb.CreatePostResponse, error) {
	postData := request.GetPost()
	data := &post{
		AuthorID: postData.GetAuthorId(),
		Title:    postData.GetTitle(),
		Content:  postData.GetContent(),
	}
	rsp, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error %s", err.Error()))
	}
	fmt.Println("Inserted: ", rsp.InsertedID)
	oid, ok := rsp.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Internal error")
	}
	return &postpb.CreatePostResponse{Post: &postpb.Post{
		Id:       oid.Hex(),
		AuthorId: postData.GetAuthorId(),
		Title:    postData.GetTitle(),
		Content:  postData.GetContent(),
	}}, nil
}

func (s *server) ReadPost(ctx context.Context, request *postpb.ReadPostRequest) (*postpb.ReadPostResponse, error) {
	postID := request.GetPostId()
	oid, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error %s", err.Error()))
	}
	data := &post{}
	filter := bson.M{"_id": oid}
	err = collection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error %s", err.Error()))
	}
	return &postpb.ReadPostResponse{Post: &postpb.Post{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Title:    data.Title,
		Content:  data.Content,
	}}, nil
}

func (s *server) ListPost(request *postpb.ListPostRequest, stream postpb.PostService_ListPostServer) error {
	cursor, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Internal error %s", err.Error()))
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		data := &post{}
		err := cursor.Decode(&data)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Error while decoding data from Mongo DB %s", err.Error()))

		}
		stream.Send(&postpb.ListPostResponse{Post: &postpb.Post{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		}})
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown server error %s", err.Error()))
	}
	return nil
}

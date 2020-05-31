package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("started blog server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	go func(){
		fmt.Println("starting server")
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("exit...")
	lis.Close()
}

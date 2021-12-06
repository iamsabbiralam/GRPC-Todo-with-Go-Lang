package main

import (
	"log"
	"net"
	
	"grpc-todo/server/todo"
	
	"google.golang.org/grpc"
	tdpb "grpc-todo/proto/todo"
)

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	grpcServer := grpc.NewServer()
	s := todo.Server{}

	tdpb.RegisterTodoServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %s", err)
	}
}
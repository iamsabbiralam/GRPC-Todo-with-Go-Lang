package main

import (
	"grpc-todo/client/todo"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}

	todoClient := todo.NewClient(conn)
	res, err := todoClient.GetTodo(3);
	if err != nil {
		log.Fatalf("error while calling GetTodo: %s", err)
	}
	
	log.Printf("Response ID: %#v", res)
	if err := todoClient.GetTodos(); err != nil {
		log.Fatalf("error while calling GetTodo: %s", err)
	}
}
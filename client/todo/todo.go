package todo

import (
	"context"
	tdpb "grpc-todo/proto/todo"
	"io"
	"log"

	"google.golang.org/grpc"
)

type Client struct {
	client	tdpb.TodoServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client {
	return Client{
		client: tdpb.NewTodoServiceClient(conn),
	}
}

func (c *Client) GetTodo(id int64) (*tdpb.GetTodoResponse, error) {
	return c.client.GetTodo(context.Background(), &tdpb.GetTodoRequest{
		ID: id,
	})
}

func (c *Client) GetTodos() error {
	stream, err := c.client.GetTodos(context.Background(), &tdpb.GetTodosRequest{})
	if err != nil {
		return err
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		log.Printf("Received ID: %d", res.GetID())	
	}
}
package todo

import (
	"context"
	"log"
	"time"

	tdpb "grpc-todo/proto/todo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {}

type Todo struct {
	ID	int64
	Title	string
	Description	string
}

var todos = []Todo {
	{
		ID : 1,
		Title: "This is Title 1",
		Description: "This is Description 1",
	},
	{
		ID : 2,
		Title: "This is Title 2",
		Description: "This is Description 2",
	},
	{
		ID : 3,
		Title: "This is Title 3",
		Description: "This is Description 3",
	},
}

func (s *Server) GetTodo(ctx context.Context, req *tdpb.GetTodoRequest) (*tdpb.GetTodoResponse, error) {
	log.Printf("Todo ID: %d", req.GetID())
	var todo Todo
	for _, value := range todos {
		if value.ID == req.GetID() {
			todo = value
			break
		}
	}

	if todo.ID == 0 {
		return &tdpb.GetTodoResponse{}, status.Errorf(codes.NotFound, "Invalid ID")
	}
	return &tdpb.GetTodoResponse{
		ID: todo.ID,
		Title: todo.Title,
		Description: todo.Description,
	}, nil 
}

func (s *Server) GetTodos(req *tdpb.GetTodosRequest, stream tdpb.TodoService_GetTodosServer) error {
	for _, value := range todos {
		err := stream.Send(&tdpb.GetTodoResponse{
			ID: value.ID,
			Title: value.Title,
			Description: value.Description,
		})
		if err != nil {
			return status.Error(codes.NotFound, "failed to send Todo")
		}
		time.Sleep(time.Second * 3)
	}
	return nil
}
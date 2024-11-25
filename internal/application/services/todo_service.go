package service

import (
	"context"

	"github.com/ehrichpeter/go-todo/internal/domain/todo"
)

type TodoService struct {
	repo *todo.TodoRepository
}

func NewTodoService(repo *todo.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetAll(ctx context.Context) ([]todo.Todo, error) {
	todos := []todo.Todo{
		todo.NewTodo("My First Todo", "Yes, it's a new todo"),
	}
	return todos, nil
}

func (s *TodoService) Create
package todo

import "context"

type TodoService interface {
	GetAll(ctx context.Context) ([]Todo, error)
	GetById(ctx context.Context, id string) (*Todo, error)
	Create(ctx context.Context, todo *Todo) error
	Delete(ctx context.Context, id string) error
}

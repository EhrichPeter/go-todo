package todo

import "context"

type TodoRepository interface {
	FindAll(ctx context.Context) ([]Todo, error)
	FindByID(ctx context.Context, id string) (*Todo, error)
	Save(ctx context.Context, todo *Todo) error
	Delete(ctx context.Context, id string) error
}

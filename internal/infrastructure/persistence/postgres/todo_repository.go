package postgres

import (
	"context"
	"database/sql"

	"github.com/ehrichpeter/go-todo/internal/domain/todo"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(ctx context.Context, t *todo.Todo) error {
	query := "INSERT INTO todos (id, title, description, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, t.ID, t.Title, t.Description, t.Completed, t.CreatedAt, t.UpdatedAt)
	return err
}

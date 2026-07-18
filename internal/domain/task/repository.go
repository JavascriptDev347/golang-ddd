package task

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Save(ctx context.Context, t *Task) error
	FindByID(ctx context.Context, id uuid.UUID) (*Task, error)
	FindAll(ctx context.Context) ([]*Task, error)
	Update(ctx context.Context, t *Task) error
}

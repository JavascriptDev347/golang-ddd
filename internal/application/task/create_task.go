package task

import (
	"context"

	domain "github.com/JavascriptDev347/golang-ddd/internal/domain/task"
)

type CreateTaskInput struct {
	Title string

	Description string
}

type CreateTaskUseCase struct {
	repo domain.Repository
}

func NewCreateTaskUseCase(repo domain.Repository) *CreateTaskUseCase {
	return &CreateTaskUseCase{repo: repo}
}

// Execute method to create a new task
func (uc *CreateTaskUseCase) Execute(ctx context.Context, in CreateTaskInput) (*domain.Task, error) {
	t, err := domain.NewTask(in.Title, in.Description)
	if err != nil {
		return nil, err
	}
	if err := uc.repo.Save(ctx, t); err != nil {
		return nil, err
	}
	return t, nil
}

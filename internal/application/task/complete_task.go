package task

import (
	"context"

	domain "github.com/JavascriptDev347/golang-ddd/internal/domain/task"
	"github.com/google/uuid"
)

type CompleteTaskUseCase struct {
	repo domain.Repository
}

func NewCompleteTaskUseCase(repo domain.Repository) *CompleteTaskUseCase {
	return &CompleteTaskUseCase{repo: repo}
}

func (uc *CompleteTaskUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	t, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if err := t.Complete(); err != nil {
		return err
	}
	return uc.repo.Update(ctx, t)
}

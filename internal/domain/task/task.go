package task

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyTitle      = errors.New("task title cannot be empty")
	ErrAlreadyComplete = errors.New("task is already completed")
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusCompleted Status = "completed"
)

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Status      Status
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(title, description string) (*Task, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}
	return &Task{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Status:      StatusPending,
		CreatedAt:   time.Now(),
	}, nil
}

func (t *Task) Complete() error {
	if t.Status == StatusCompleted {
		return ErrAlreadyComplete
	}
	t.Status = StatusCompleted
	now := time.Now()
	t.CompletedAt = &now
	return nil
}

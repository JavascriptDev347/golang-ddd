package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	domain "github.com/JavascriptDev347/golang-ddd/internal/domain/task"
)

type TaskRepository struct {
	pool *pgxpool.Pool
}

func NewTaskRepository(pool *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{pool: pool}
}

func (r *TaskRepository) Save(ctx context.Context, t *domain.Task) error {
	_, err := r.pool.Exec(ctx,
		`INSERT INTO tasks (id, title, description, status, created_at)
		 VALUES ($1, $2, $3, $4, $5)`,
		t.ID, t.Title, t.Description, t.Status, t.CreatedAt,
	)
	return err
}
func (r *TaskRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	row := r.pool.QueryRow(ctx,
		`SELECT id, title, description, status, created_at, completed_at
		 FROM tasks WHERE id = $1`, id)

	var t domain.Task
	err := row.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.CompletedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrEmptyTitle // amalda o'z ErrNotFound xatoingni yarat
	}
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TaskRepository) Update(ctx context.Context, t *domain.Task) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE tasks SET status=$1, completed_at=$2 WHERE id=$3`,
		t.Status, t.CompletedAt, t.ID,
	)
	return err
}
func (r *TaskRepository) FindAll(ctx context.Context) ([]*domain.Task, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, title, description, status, created_at, completed_at FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		var t domain.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.CompletedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, nil
}

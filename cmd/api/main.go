package main

import (
	"context"
	"log"

	appTask "github.com/JavascriptDev347/golang-ddd/internal/application/task"
	"github.com/JavascriptDev347/golang-ddd/internal/infrastructure/postgres"
	httpi "github.com/JavascriptDev347/golang-ddd/internal/interfaces/http"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgres://postgres:postgres@localhost:5432/go-ddd")
	if err != nil {
		log.Fatalf("db connection failed: %v", err)
	}
	defer pool.Close()

	repo := postgres.NewTaskRepository(pool)

	createUC := appTask.NewCreateTaskUseCase(repo)
	completeUC := appTask.NewCompleteTaskUseCase(repo)

	handler := httpi.NewTaskHandler(createUC, completeUC)
	router := httpi.NewRouter(handler)

	router.Run(":8080")
}

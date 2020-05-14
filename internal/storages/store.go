package storages

import (
	"context"

	models "../models"
)

type Store interface {
	GetTaskList(ctx context.Context) (models.TaskList, error)
	CreateTask(ctx context.Context, task models.Task) (models.Task, error)
	DeleteTask(ctx context.Context, id string) error
}

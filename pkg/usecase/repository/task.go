package repository

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
)

// Task is interface of repository
type Task interface {
	Get(ctx context.Context, where *model.TaskWhereInput) (*model.Task, error)
	List(ctx context.Context) ([]*model.Task, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskWhereInput) (*model.TaskConnection, error)
	Create(ctx context.Context, input model.CreateTaskInput) (*model.Task, error)
	Update(ctx context.Context, input model.UpdateTaskInput) (*model.Task, error)
	Delete(ctx context.Context, input model.DeleteTaskInput) (*model.DeleteTaskPayload, error)
	Undelete(ctx context.Context, input model.UndeleteTaskInput) (*model.UndeleteTaskPayload, error)
}

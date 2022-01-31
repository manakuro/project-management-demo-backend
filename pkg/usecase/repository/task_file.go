package repository

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
)

// TaskFile is interface of repository
type TaskFile interface {
	Get(ctx context.Context, where *model.TaskFileWhereInput) (*model.TaskFile, error)
	List(ctx context.Context) ([]*model.TaskFile, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskFileWhereInput, requestedFields []string) (*model.TaskFileConnection, error)
	Create(ctx context.Context, input model.CreateTaskFileInput) (*model.TaskFile, error)
	Update(ctx context.Context, input model.UpdateTaskFileInput) (*model.TaskFile, error)
}

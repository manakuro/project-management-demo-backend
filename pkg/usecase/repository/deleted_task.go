package repository

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
)

// DeletedTask is interface of repository
type DeletedTask interface {
	Get(ctx context.Context, where *model.DeletedTaskWhereInput) (*model.DeletedTask, error)
	List(ctx context.Context) ([]*model.DeletedTask, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.DeletedTaskWhereInput) (*model.DeletedTaskConnection, error)
	Create(ctx context.Context, input model.CreateDeletedTaskInput) (*model.DeletedTask, error)
	Update(ctx context.Context, input model.UpdateDeletedTaskInput) (*model.DeletedTask, error)
	Delete(ctx context.Context, input model.DeleteDeletedTaskInput) (*model.DeletedTask, error)
}

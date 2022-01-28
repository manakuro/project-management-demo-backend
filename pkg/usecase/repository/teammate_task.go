package repository

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
)

// TeammateTask is interface of repository
type TeammateTask interface {
	Get(ctx context.Context, where *model.TeammateTaskWhereInput) (*model.TeammateTask, error)
	List(ctx context.Context) ([]*model.TeammateTask, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskWhereInput, requestedFields []string) (*model.TeammateTaskConnection, error)
	Create(ctx context.Context, input model.CreateTeammateTaskInput) (*model.TeammateTask, error)
	Update(ctx context.Context, input model.UpdateTeammateTaskInput) (*model.TeammateTask, error)
	Delete(ctx context.Context, input model.DeleteTeammateTaskInput) (*model.TeammateTask, error)
}

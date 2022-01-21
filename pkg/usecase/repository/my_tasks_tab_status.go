package repository

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
)

// MyTasksTabStatus is interface of repository
type MyTasksTabStatus interface {
	Get(ctx context.Context, where *model.MyTasksTabStatusWhereInput) (*model.MyTasksTabStatus, error)
	List(ctx context.Context) ([]*model.MyTasksTabStatus, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.MyTasksTabStatusWhereInput, requestedFields []string) (*model.MyTasksTabStatusConnection, error)
	Create(ctx context.Context, input model.CreateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error)
	Update(ctx context.Context, input model.UpdateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error)
}

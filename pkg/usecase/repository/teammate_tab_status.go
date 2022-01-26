package repository

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
)

// TeammateTabStatus is interface of repository
type TeammateTabStatus interface {
	Get(ctx context.Context, where *model.TeammateTabStatusWhereInput) (*model.TeammateTabStatus, error)
	List(ctx context.Context) ([]*model.TeammateTabStatus, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTabStatusWhereInput, requestedFields []string) (*model.TeammateTabStatusConnection, error)
	Create(ctx context.Context, input model.CreateTeammateTabStatusInput) (*model.TeammateTabStatus, error)
	Update(ctx context.Context, input model.UpdateTeammateTabStatusInput) (*model.TeammateTabStatus, error)
}

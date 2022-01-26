package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TeammateTabStatus is an interface of controller
type TeammateTabStatus interface {
	Get(ctx context.Context, where *model.TeammateTabStatusWhereInput) (*model.TeammateTabStatus, error)
	List(ctx context.Context) ([]*model.TeammateTabStatus, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTabStatusWhereInput, requestedFields []string) (*model.TeammateTabStatusConnection, error)
	Create(ctx context.Context, input model.CreateTeammateTabStatusInput) (*model.TeammateTabStatus, error)
	Update(ctx context.Context, input model.UpdateTeammateTabStatusInput) (*model.TeammateTabStatus, error)
}

type teammateTabStatusController struct {
	teammateTabStatusUsecase usecase.TeammateTabStatus
}

// NewTeammateTabStatusController generates teammateTabStatus controller
func NewTeammateTabStatusController(pt usecase.TeammateTabStatus) TeammateTabStatus {
	return &teammateTabStatusController{
		teammateTabStatusUsecase: pt,
	}
}

func (c *teammateTabStatusController) Get(ctx context.Context, where *model.TeammateTabStatusWhereInput) (*model.TeammateTabStatus, error) {
	return c.teammateTabStatusUsecase.Get(ctx, where)
}

func (c *teammateTabStatusController) List(ctx context.Context) ([]*model.TeammateTabStatus, error) {
	return c.teammateTabStatusUsecase.List(ctx)
}

func (c *teammateTabStatusController) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTabStatusWhereInput, requestedFields []string) (*model.TeammateTabStatusConnection, error) {
	return c.teammateTabStatusUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (c *teammateTabStatusController) Create(ctx context.Context, input model.CreateTeammateTabStatusInput) (*model.TeammateTabStatus, error) {
	return c.teammateTabStatusUsecase.Create(ctx, input)
}

func (c *teammateTabStatusController) Update(ctx context.Context, input model.UpdateTeammateTabStatusInput) (*model.TeammateTabStatus, error) {
	return c.teammateTabStatusUsecase.Update(ctx, input)
}

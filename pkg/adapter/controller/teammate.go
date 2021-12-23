package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// Teammate is an interface of controller
type Teammate interface {
	Get(ctx context.Context, id model.ID) (*model.Teammate, error)
	List(ctx context.Context) ([]*model.Teammate, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateWhereInput) (*model.TeammateConnection, error)
	Create(ctx context.Context, input model.CreateTeammateInput) (*model.Teammate, error)
	Update(ctx context.Context, input model.UpdateTeammateInput) (*model.Teammate, error)
}

type teammate struct {
	teammateUsecase usecase.Teammate
}

// NewTeammateController generates teammate controller
func NewTeammateController(tu usecase.Teammate) Teammate {
	return &teammate{
		teammateUsecase: tu,
	}
}

func (t *teammate) Get(ctx context.Context, id model.ID) (*model.Teammate, error) {
	return t.teammateUsecase.Get(ctx, id)
}

func (t *teammate) List(ctx context.Context) ([]*model.Teammate, error) {
	return t.teammateUsecase.List(ctx)
}

func (t *teammate) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateWhereInput) (*model.TeammateConnection, error) {
	return t.teammateUsecase.ListWithPagination(ctx, after, first, before, last, where)
}

func (t *teammate) Create(ctx context.Context, input model.CreateTeammateInput) (*model.Teammate, error) {
	return t.teammateUsecase.Create(ctx, input)
}

func (t *teammate) Update(ctx context.Context, input model.UpdateTeammateInput) (*model.Teammate, error) {
	return t.teammateUsecase.Update(ctx, input)
}

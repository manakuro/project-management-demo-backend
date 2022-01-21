package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TeammateTaskColumn is an interface of controller
type TeammateTaskColumn interface {
	Get(ctx context.Context, where *model.TeammateTaskColumnWhereInput) (*model.TeammateTaskColumn, error)
	List(ctx context.Context) ([]*model.TeammateTaskColumn, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskColumnWhereInput, requestedFields []string) (*model.TeammateTaskColumnConnection, error)
	Create(ctx context.Context, input model.CreateTeammateTaskColumnInput) (*model.TeammateTaskColumn, error)
	Update(ctx context.Context, input model.UpdateTeammateTaskColumnInput) (*model.TeammateTaskColumn, error)
}

type teammateTaskColumn struct {
	teammateTaskColumnUsecase usecase.TeammateTaskColumn
}

// NewTeammateTaskColumnController generates teammateTaskColumn controller
func NewTeammateTaskColumnController(pt usecase.TeammateTaskColumn) TeammateTaskColumn {
	return &teammateTaskColumn{
		teammateTaskColumnUsecase: pt,
	}
}

func (s *teammateTaskColumn) Get(ctx context.Context, where *model.TeammateTaskColumnWhereInput) (*model.TeammateTaskColumn, error) {
	return s.teammateTaskColumnUsecase.Get(ctx, where)
}

func (s *teammateTaskColumn) List(ctx context.Context) ([]*model.TeammateTaskColumn, error) {
	return s.teammateTaskColumnUsecase.List(ctx)
}

func (s *teammateTaskColumn) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskColumnWhereInput, requestedFields []string) (*model.TeammateTaskColumnConnection, error) {
	return s.teammateTaskColumnUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (s *teammateTaskColumn) Create(ctx context.Context, input model.CreateTeammateTaskColumnInput) (*model.TeammateTaskColumn, error) {
	return s.teammateTaskColumnUsecase.Create(ctx, input)
}

func (s *teammateTaskColumn) Update(ctx context.Context, input model.UpdateTeammateTaskColumnInput) (*model.TeammateTaskColumn, error) {
	return s.teammateTaskColumnUsecase.Update(ctx, input)
}

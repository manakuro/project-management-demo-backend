package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// Color is an interface of controller
type Color interface {
	Get(ctx context.Context, id model.ID) (*model.Color, error)
	List(ctx context.Context) ([]*model.Color, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ColorWhereInput) (*model.ColorConnection, error)
	Create(ctx context.Context, input model.CreateColorInput) (*model.Color, error)
	Update(ctx context.Context, input model.UpdateColorInput) (*model.Color, error)
}

type color struct {
	colorUsecase usecase.Color
}

// NewColorController generates color controller
func NewColorController(tu usecase.Color) Color {
	return &color{
		colorUsecase: tu,
	}
}

func (t *color) Get(ctx context.Context, id model.ID) (*model.Color, error) {
	return t.colorUsecase.Get(ctx, id)
}

func (t *color) List(ctx context.Context) ([]*model.Color, error) {
	return t.colorUsecase.List(ctx)
}

func (t *color) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ColorWhereInput) (*model.ColorConnection, error) {
	return t.colorUsecase.ListWithPagination(ctx, after, first, before, last, where)
}

func (t *color) Create(ctx context.Context, input model.CreateColorInput) (*model.Color, error) {
	return t.colorUsecase.Create(ctx, input)
}

func (t *color) Update(ctx context.Context, input model.UpdateColorInput) (*model.Color, error) {
	return t.colorUsecase.Update(ctx, input)
}

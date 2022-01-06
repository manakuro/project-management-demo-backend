package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type color struct {
	colorRepository repository.Color
}

// Color is an interface of test user
type Color interface {
	Get(ctx context.Context, id model.ID) (*model.Color, error)
	List(ctx context.Context) ([]*model.Color, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ColorWhereInput) (*model.ColorConnection, error)
	Create(ctx context.Context, input model.CreateColorInput) (*model.Color, error)
	Update(ctx context.Context, input model.UpdateColorInput) (*model.Color, error)
}

// NewColorUsecase generates test user repository
func NewColorUsecase(r repository.Color) Color {
	return &color{colorRepository: r}
}

func (t *color) Get(ctx context.Context, id model.ID) (*model.Color, error) {
	return t.colorRepository.Get(ctx, id)
}

func (t *color) List(ctx context.Context) ([]*model.Color, error) {
	return t.colorRepository.List(ctx)
}

func (t *color) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ColorWhereInput) (*model.ColorConnection, error) {
	return t.colorRepository.ListWithPagination(ctx, after, first, before, last, where)
}

func (t *color) Create(ctx context.Context, input model.CreateColorInput) (*model.Color, error) {
	return t.colorRepository.Create(ctx, input)
}

func (t *color) Update(ctx context.Context, input model.UpdateColorInput) (*model.Color, error) {
	return t.colorRepository.Update(ctx, input)
}

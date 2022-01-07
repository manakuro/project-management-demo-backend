package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// Icon is an interface of controller
type Icon interface {
	Get(ctx context.Context, id model.ID) (*model.Icon, error)
	List(ctx context.Context) ([]*model.Icon, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.IconWhereInput) (*model.IconConnection, error)
	Create(ctx context.Context, input model.CreateIconInput) (*model.Icon, error)
	Update(ctx context.Context, input model.UpdateIconInput) (*model.Icon, error)
}

type icon struct {
	iconUsecase usecase.Icon
}

// NewIconController generates icon controller
func NewIconController(tu usecase.Icon) Icon {
	return &icon{
		iconUsecase: tu,
	}
}

func (t *icon) Get(ctx context.Context, id model.ID) (*model.Icon, error) {
	return t.iconUsecase.Get(ctx, id)
}

func (t *icon) List(ctx context.Context) ([]*model.Icon, error) {
	return t.iconUsecase.List(ctx)
}

func (t *icon) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.IconWhereInput) (*model.IconConnection, error) {
	return t.iconUsecase.ListWithPagination(ctx, after, first, before, last, where)
}

func (t *icon) Create(ctx context.Context, input model.CreateIconInput) (*model.Icon, error) {
	return t.iconUsecase.Create(ctx, input)
}

func (t *icon) Update(ctx context.Context, input model.UpdateIconInput) (*model.Icon, error) {
	return t.iconUsecase.Update(ctx, input)
}

package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type icon struct {
	iconRepository repository.Icon
}

// Icon is an interface of test user
type Icon interface {
	Get(ctx context.Context, id model.ID) (*model.Icon, error)
	List(ctx context.Context) ([]*model.Icon, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.IconWhereInput) (*model.IconConnection, error)
	Create(ctx context.Context, input model.CreateIconInput) (*model.Icon, error)
	Update(ctx context.Context, input model.UpdateIconInput) (*model.Icon, error)
}

// NewIconUsecase generates test user repository
func NewIconUsecase(r repository.Icon) Icon {
	return &icon{iconRepository: r}
}

func (t *icon) Get(ctx context.Context, id model.ID) (*model.Icon, error) {
	return t.iconRepository.Get(ctx, id)
}

func (t *icon) List(ctx context.Context) ([]*model.Icon, error) {
	return t.iconRepository.List(ctx)
}

func (t *icon) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.IconWhereInput) (*model.IconConnection, error) {
	return t.iconRepository.ListWithPagination(ctx, after, first, before, last, where)
}

func (t *icon) Create(ctx context.Context, input model.CreateIconInput) (*model.Icon, error) {
	return t.iconRepository.Create(ctx, input)
}

func (t *icon) Update(ctx context.Context, input model.UpdateIconInput) (*model.Icon, error) {
	return t.iconRepository.Update(ctx, input)
}

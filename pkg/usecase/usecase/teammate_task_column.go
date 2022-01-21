package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type teammateTaskColumn struct {
	teammateTaskColumnRepository repository.TeammateTaskColumn
}

// TeammateTaskColumn is an interface of test user
type TeammateTaskColumn interface {
	Get(ctx context.Context, where *model.TeammateTaskColumnWhereInput) (*model.TeammateTaskColumn, error)
	List(ctx context.Context) ([]*model.TeammateTaskColumn, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskColumnWhereInput, requestedFields []string) (*model.TeammateTaskColumnConnection, error)
	Create(ctx context.Context, input model.CreateTeammateTaskColumnInput) (*model.TeammateTaskColumn, error)
	Update(ctx context.Context, input model.UpdateTeammateTaskColumnInput) (*model.TeammateTaskColumn, error)
}

// NewTeammateTaskColumnUsecase generates test user repository
func NewTeammateTaskColumnUsecase(r repository.TeammateTaskColumn) TeammateTaskColumn {
	return &teammateTaskColumn{teammateTaskColumnRepository: r}
}

func (c *teammateTaskColumn) Get(ctx context.Context, where *model.TeammateTaskColumnWhereInput) (*model.TeammateTaskColumn, error) {
	return c.teammateTaskColumnRepository.Get(ctx, where)
}

func (c *teammateTaskColumn) List(ctx context.Context) ([]*model.TeammateTaskColumn, error) {
	return c.teammateTaskColumnRepository.List(ctx)
}

func (c *teammateTaskColumn) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskColumnWhereInput, requestedFields []string) (*model.TeammateTaskColumnConnection, error) {
	return c.teammateTaskColumnRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (c *teammateTaskColumn) Create(ctx context.Context, input model.CreateTeammateTaskColumnInput) (*model.TeammateTaskColumn, error) {
	return c.teammateTaskColumnRepository.Create(ctx, input)
}

func (c *teammateTaskColumn) Update(ctx context.Context, input model.UpdateTeammateTaskColumnInput) (*model.TeammateTaskColumn, error) {
	return c.teammateTaskColumnRepository.Update(ctx, input)
}

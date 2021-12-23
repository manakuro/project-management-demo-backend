package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type teammate struct {
	teammateRepository repository.Teammate
}

// Teammate is an interface of test user
type Teammate interface {
	Get(ctx context.Context, id model.ID) (*model.Teammate, error)
	List(ctx context.Context) ([]*model.Teammate, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateWhereInput) (*model.TeammateConnection, error)
	Create(ctx context.Context, input model.CreateTeammateInput) (*model.Teammate, error)
	Update(ctx context.Context, input model.UpdateTeammateInput) (*model.Teammate, error)
}

// NewTeammateUsecase generates test user repository
func NewTeammateUsecase(r repository.Teammate) Teammate {
	return &teammate{teammateRepository: r}
}

func (t *teammate) Get(ctx context.Context, id model.ID) (*model.Teammate, error) {
	return t.teammateRepository.Get(ctx, id)
}

func (t *teammate) List(ctx context.Context) ([]*model.Teammate, error) {
	return t.teammateRepository.List(ctx)
}

func (t *teammate) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateWhereInput) (*model.TeammateConnection, error) {
	return t.teammateRepository.ListWithPagination(ctx, after, first, before, last, where)
}

func (t *teammate) Create(ctx context.Context, input model.CreateTeammateInput) (*model.Teammate, error) {
	return t.teammateRepository.Create(ctx, input)
}

func (t *teammate) Update(ctx context.Context, input model.UpdateTeammateInput) (*model.Teammate, error) {
	return t.teammateRepository.Update(ctx, input)
}

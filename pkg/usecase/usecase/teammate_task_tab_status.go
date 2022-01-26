package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type teammateTaskTabStatusUsecase struct {
	teammateTaskTabStatusRepository repository.TeammateTaskTabStatus
}

// TeammateTaskTabStatus is an interface of test user
type TeammateTaskTabStatus interface {
	Get(ctx context.Context, where *model.TeammateTaskTabStatusWhereInput) (*model.TeammateTaskTabStatus, error)
	List(ctx context.Context) ([]*model.TeammateTaskTabStatus, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskTabStatusWhereInput, requestedFields []string) (*model.TeammateTaskTabStatusConnection, error)
	Create(ctx context.Context, input model.CreateTeammateTaskTabStatusInput) (*model.TeammateTaskTabStatus, error)
	Update(ctx context.Context, input model.UpdateTeammateTaskTabStatusInput) (*model.TeammateTaskTabStatus, error)
}

// NewTeammateTaskTabStatusUsecase generates test user repository
func NewTeammateTaskTabStatusUsecase(r repository.TeammateTaskTabStatus) TeammateTaskTabStatus {
	return &teammateTaskTabStatusUsecase{teammateTaskTabStatusRepository: r}
}

func (u *teammateTaskTabStatusUsecase) Get(ctx context.Context, where *model.TeammateTaskTabStatusWhereInput) (*model.TeammateTaskTabStatus, error) {
	return u.teammateTaskTabStatusRepository.Get(ctx, where)
}

func (u *teammateTaskTabStatusUsecase) List(ctx context.Context) ([]*model.TeammateTaskTabStatus, error) {
	return u.teammateTaskTabStatusRepository.List(ctx)
}

func (u *teammateTaskTabStatusUsecase) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskTabStatusWhereInput, requestedFields []string) (*model.TeammateTaskTabStatusConnection, error) {
	return u.teammateTaskTabStatusRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (u *teammateTaskTabStatusUsecase) Create(ctx context.Context, input model.CreateTeammateTaskTabStatusInput) (*model.TeammateTaskTabStatus, error) {
	return u.teammateTaskTabStatusRepository.Create(ctx, input)
}

func (u *teammateTaskTabStatusUsecase) Update(ctx context.Context, input model.UpdateTeammateTaskTabStatusInput) (*model.TeammateTaskTabStatus, error) {
	return u.teammateTaskTabStatusRepository.Update(ctx, input)
}

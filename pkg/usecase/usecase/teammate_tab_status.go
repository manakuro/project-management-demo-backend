package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type teammateTabStatusUsecase struct {
	teammateTabStatusRepository repository.TeammateTabStatus
}

// TeammateTabStatus is an interface of test user
type TeammateTabStatus interface {
	Get(ctx context.Context, where *model.TeammateTabStatusWhereInput) (*model.TeammateTabStatus, error)
	List(ctx context.Context) ([]*model.TeammateTabStatus, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTabStatusWhereInput, requestedFields []string) (*model.TeammateTabStatusConnection, error)
	Create(ctx context.Context, input model.CreateTeammateTabStatusInput) (*model.TeammateTabStatus, error)
	Update(ctx context.Context, input model.UpdateTeammateTabStatusInput) (*model.TeammateTabStatus, error)
}

// NewTeammateTabStatusUsecase generates test user repository
func NewTeammateTabStatusUsecase(r repository.TeammateTabStatus) TeammateTabStatus {
	return &teammateTabStatusUsecase{teammateTabStatusRepository: r}
}

func (u *teammateTabStatusUsecase) Get(ctx context.Context, where *model.TeammateTabStatusWhereInput) (*model.TeammateTabStatus, error) {
	return u.teammateTabStatusRepository.Get(ctx, where)
}

func (u *teammateTabStatusUsecase) List(ctx context.Context) ([]*model.TeammateTabStatus, error) {
	return u.teammateTabStatusRepository.List(ctx)
}

func (u *teammateTabStatusUsecase) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTabStatusWhereInput, requestedFields []string) (*model.TeammateTabStatusConnection, error) {
	return u.teammateTabStatusRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (u *teammateTabStatusUsecase) Create(ctx context.Context, input model.CreateTeammateTabStatusInput) (*model.TeammateTabStatus, error) {
	return u.teammateTabStatusRepository.Create(ctx, input)
}

func (u *teammateTabStatusUsecase) Update(ctx context.Context, input model.UpdateTeammateTabStatusInput) (*model.TeammateTabStatus, error) {
	return u.teammateTabStatusRepository.Update(ctx, input)
}

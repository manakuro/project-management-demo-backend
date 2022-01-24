package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type myTasksTabStatusUsecase struct {
	myTasksTabStatusRepository repository.MyTasksTabStatus
}

// MyTasksTabStatus is an interface of test user
type MyTasksTabStatus interface {
	Get(ctx context.Context, where *model.MyTasksTabStatusWhereInput) (*model.MyTasksTabStatus, error)
	List(ctx context.Context) ([]*model.MyTasksTabStatus, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.MyTasksTabStatusWhereInput, requestedFields []string) (*model.MyTasksTabStatusConnection, error)
	Create(ctx context.Context, input model.CreateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error)
	Update(ctx context.Context, input model.UpdateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error)
}

// NewMyTasksTabStatusUsecase generates test user repository
func NewMyTasksTabStatusUsecase(r repository.MyTasksTabStatus) MyTasksTabStatus {
	return &myTasksTabStatusUsecase{myTasksTabStatusRepository: r}
}

func (u *myTasksTabStatusUsecase) Get(ctx context.Context, where *model.MyTasksTabStatusWhereInput) (*model.MyTasksTabStatus, error) {
	return u.myTasksTabStatusRepository.Get(ctx, where)
}

func (u *myTasksTabStatusUsecase) List(ctx context.Context) ([]*model.MyTasksTabStatus, error) {
	return u.myTasksTabStatusRepository.List(ctx)
}

func (u *myTasksTabStatusUsecase) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.MyTasksTabStatusWhereInput, requestedFields []string) (*model.MyTasksTabStatusConnection, error) {
	return u.myTasksTabStatusRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (u *myTasksTabStatusUsecase) Create(ctx context.Context, input model.CreateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error) {
	return u.myTasksTabStatusRepository.Create(ctx, input)
}

func (u *myTasksTabStatusUsecase) Update(ctx context.Context, input model.UpdateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error) {
	return u.myTasksTabStatusRepository.Update(ctx, input)
}

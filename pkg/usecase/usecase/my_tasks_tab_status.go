package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type myTasksTabStatus struct {
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
	return &myTasksTabStatus{myTasksTabStatusRepository: r}
}

func (s *myTasksTabStatus) Get(ctx context.Context, where *model.MyTasksTabStatusWhereInput) (*model.MyTasksTabStatus, error) {
	return s.myTasksTabStatusRepository.Get(ctx, where)
}

func (s *myTasksTabStatus) List(ctx context.Context) ([]*model.MyTasksTabStatus, error) {
	return s.myTasksTabStatusRepository.List(ctx)
}

func (s *myTasksTabStatus) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.MyTasksTabStatusWhereInput, requestedFields []string) (*model.MyTasksTabStatusConnection, error) {
	return s.myTasksTabStatusRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (s *myTasksTabStatus) Create(ctx context.Context, input model.CreateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error) {
	return s.myTasksTabStatusRepository.Create(ctx, input)
}

func (s *myTasksTabStatus) Update(ctx context.Context, input model.UpdateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error) {
	return s.myTasksTabStatusRepository.Update(ctx, input)
}

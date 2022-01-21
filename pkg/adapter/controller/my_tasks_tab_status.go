package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// MyTasksTabStatus is an interface of controller
type MyTasksTabStatus interface {
	Get(ctx context.Context, where *model.MyTasksTabStatusWhereInput) (*model.MyTasksTabStatus, error)
	List(ctx context.Context) ([]*model.MyTasksTabStatus, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.MyTasksTabStatusWhereInput, requestedFields []string) (*model.MyTasksTabStatusConnection, error)
	Create(ctx context.Context, input model.CreateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error)
	Update(ctx context.Context, input model.UpdateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error)
}

type myTasksTabStatus struct {
	myTasksTabStatusUsecase usecase.MyTasksTabStatus
}

// NewMyTasksTabStatusController generates myTasksTabStatus controller
func NewMyTasksTabStatusController(pt usecase.MyTasksTabStatus) MyTasksTabStatus {
	return &myTasksTabStatus{
		myTasksTabStatusUsecase: pt,
	}
}

func (s *myTasksTabStatus) Get(ctx context.Context, where *model.MyTasksTabStatusWhereInput) (*model.MyTasksTabStatus, error) {
	return s.myTasksTabStatusUsecase.Get(ctx, where)
}

func (s *myTasksTabStatus) List(ctx context.Context) ([]*model.MyTasksTabStatus, error) {
	return s.myTasksTabStatusUsecase.List(ctx)
}

func (s *myTasksTabStatus) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.MyTasksTabStatusWhereInput, requestedFields []string) (*model.MyTasksTabStatusConnection, error) {
	return s.myTasksTabStatusUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (s *myTasksTabStatus) Create(ctx context.Context, input model.CreateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error) {
	return s.myTasksTabStatusUsecase.Create(ctx, input)
}

func (s *myTasksTabStatus) Update(ctx context.Context, input model.UpdateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error) {
	return s.myTasksTabStatusUsecase.Update(ctx, input)
}

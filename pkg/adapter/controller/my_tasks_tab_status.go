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

type myTasksTabStatusController struct {
	myTasksTabStatusUsecase usecase.MyTasksTabStatus
}

// NewMyTasksTabStatusController generates myTasksTabStatus controller
func NewMyTasksTabStatusController(pt usecase.MyTasksTabStatus) MyTasksTabStatus {
	return &myTasksTabStatusController{
		myTasksTabStatusUsecase: pt,
	}
}

func (c *myTasksTabStatusController) Get(ctx context.Context, where *model.MyTasksTabStatusWhereInput) (*model.MyTasksTabStatus, error) {
	return c.myTasksTabStatusUsecase.Get(ctx, where)
}

func (c *myTasksTabStatusController) List(ctx context.Context) ([]*model.MyTasksTabStatus, error) {
	return c.myTasksTabStatusUsecase.List(ctx)
}

func (c *myTasksTabStatusController) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.MyTasksTabStatusWhereInput, requestedFields []string) (*model.MyTasksTabStatusConnection, error) {
	return c.myTasksTabStatusUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (c *myTasksTabStatusController) Create(ctx context.Context, input model.CreateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error) {
	return c.myTasksTabStatusUsecase.Create(ctx, input)
}

func (c *myTasksTabStatusController) Update(ctx context.Context, input model.UpdateMyTasksTabStatusInput) (*model.MyTasksTabStatus, error) {
	return c.myTasksTabStatusUsecase.Update(ctx, input)
}

package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TaskColumn is an interface of controller
type TaskColumn interface {
	Get(ctx context.Context, where *model.TaskColumnWhereInput) (*model.TaskColumn, error)
	List(ctx context.Context) ([]*model.TaskColumn, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskColumnWhereInput, requestedFields []string) (*model.TaskColumnConnection, error)
	Create(ctx context.Context, input model.CreateTaskColumnInput) (*model.TaskColumn, error)
	Update(ctx context.Context, input model.UpdateTaskColumnInput) (*model.TaskColumn, error)
}

type taskColumn struct {
	taskColumnUsecase usecase.TaskColumn
}

// NewTaskColumnController generates taskColumn controller
func NewTaskColumnController(pt usecase.TaskColumn) TaskColumn {
	return &taskColumn{
		taskColumnUsecase: pt,
	}
}

func (c *taskColumn) Get(ctx context.Context, where *model.TaskColumnWhereInput) (*model.TaskColumn, error) {
	return c.taskColumnUsecase.Get(ctx, where)
}

func (c *taskColumn) List(ctx context.Context) ([]*model.TaskColumn, error) {
	return c.taskColumnUsecase.List(ctx)
}

func (c *taskColumn) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskColumnWhereInput, requestedFields []string) (*model.TaskColumnConnection, error) {
	return c.taskColumnUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (c *taskColumn) Create(ctx context.Context, input model.CreateTaskColumnInput) (*model.TaskColumn, error) {
	return c.taskColumnUsecase.Create(ctx, input)
}

func (c *taskColumn) Update(ctx context.Context, input model.UpdateTaskColumnInput) (*model.TaskColumn, error) {
	return c.taskColumnUsecase.Update(ctx, input)
}

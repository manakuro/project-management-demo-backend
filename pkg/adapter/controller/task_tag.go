package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TaskTag is an interface of controller.
type TaskTag interface {
	Get(ctx context.Context, where *model.TaskTagWhereInput) (*model.TaskTag, error)
	List(ctx context.Context, where *model.TaskTagWhereInput) ([]*model.TaskTag, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskTagWhereInput, requestedFields []string) (*model.TaskTagConnection, error)
	Create(ctx context.Context, input model.CreateTaskTagInput) (*model.TaskTag, error)
	Update(ctx context.Context, input model.UpdateTaskTagInput) (*model.TaskTag, error)
	Delete(ctx context.Context, input model.DeleteTaskTagInput) (*model.TaskTag, error)
}

type taskTagController struct {
	taskTagUsecase usecase.TaskTag
}

// NewTaskTagController generates controller.
func NewTaskTagController(u usecase.TaskTag) TaskTag {
	return &taskTagController{
		taskTagUsecase: u,
	}
}

func (c *taskTagController) Get(ctx context.Context, where *model.TaskTagWhereInput) (*model.TaskTag, error) {
	return c.taskTagUsecase.Get(ctx, where)
}

func (c *taskTagController) List(ctx context.Context, where *model.TaskTagWhereInput) ([]*model.TaskTag, error) {
	return c.taskTagUsecase.List(ctx, where)
}

func (c *taskTagController) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskTagWhereInput, requestedFields []string) (*model.TaskTagConnection, error) {
	return c.taskTagUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (c *taskTagController) Create(ctx context.Context, input model.CreateTaskTagInput) (*model.TaskTag, error) {
	return c.taskTagUsecase.Create(ctx, input)
}

func (c *taskTagController) Update(ctx context.Context, input model.UpdateTaskTagInput) (*model.TaskTag, error) {
	return c.taskTagUsecase.Update(ctx, input)
}

func (c *taskTagController) Delete(ctx context.Context, input model.DeleteTaskTagInput) (*model.TaskTag, error) {
	return c.taskTagUsecase.Delete(ctx, input)
}

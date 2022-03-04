package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// Task is an interface of controller.
type Task interface {
	Get(ctx context.Context, where *model.TaskWhereInput) (*model.Task, error)
	List(ctx context.Context) ([]*model.Task, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskWhereInput) (*model.TaskConnection, error)
	Create(ctx context.Context, input model.CreateTaskInput) (*model.Task, error)
	Update(ctx context.Context, input model.UpdateTaskInput) (*model.Task, error)
	Delete(ctx context.Context, input model.DeleteTaskInput) (*model.DeleteTaskPayload, error)
	DeleteAll(ctx context.Context, input model.DeleteAllTaskInput) (*model.DeleteAllTaskPayload, error)
	Undelete(ctx context.Context, input model.UndeleteTaskInput) (*model.UndeleteTaskPayload, error)
	UndeleteAll(ctx context.Context, input model.UndeleteAllTaskInput) (*model.UndeleteAllTaskPayload, error)
}

type taskController struct {
	taskUsecase usecase.Task
}

// NewTaskController generates task controller.
func NewTaskController(pt usecase.Task) Task {
	return &taskController{
		taskUsecase: pt,
	}
}

func (c *taskController) Get(ctx context.Context, where *model.TaskWhereInput) (*model.Task, error) {
	return c.taskUsecase.Get(ctx, where)
}

func (c *taskController) List(ctx context.Context) ([]*model.Task, error) {
	return c.taskUsecase.List(ctx)
}

func (c *taskController) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskWhereInput) (*model.TaskConnection, error) {
	return c.taskUsecase.ListWithPagination(ctx, after, first, before, last, where)
}

func (c *taskController) Create(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	return c.taskUsecase.Create(ctx, input)
}

func (c *taskController) Update(ctx context.Context, input model.UpdateTaskInput) (*model.Task, error) {
	return c.taskUsecase.Update(ctx, input)
}

func (c *taskController) Delete(ctx context.Context, input model.DeleteTaskInput) (*model.DeleteTaskPayload, error) {
	return c.taskUsecase.Delete(ctx, input)
}

func (c *taskController) DeleteAll(ctx context.Context, input model.DeleteAllTaskInput) (*model.DeleteAllTaskPayload, error) {
	return c.taskUsecase.DeleteAll(ctx, input)
}

func (c *taskController) Undelete(ctx context.Context, input model.UndeleteTaskInput) (*model.UndeleteTaskPayload, error) {
	return c.taskUsecase.Undelete(ctx, input)
}

func (c *taskController) UndeleteAll(ctx context.Context, input model.UndeleteAllTaskInput) (*model.UndeleteAllTaskPayload, error) {
	return c.taskUsecase.UndeleteAll(ctx, input)
}

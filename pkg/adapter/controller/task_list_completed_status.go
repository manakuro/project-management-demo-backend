package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TaskListCompletedStatus is an interface of controller.
type TaskListCompletedStatus interface {
	Get(ctx context.Context, where *model.TaskListCompletedStatusWhereInput) (*model.TaskListCompletedStatus, error)
	List(ctx context.Context) ([]*model.TaskListCompletedStatus, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskListCompletedStatusWhereInput, requestedFields []string) (*model.TaskListCompletedStatusConnection, error)
	Create(ctx context.Context, input model.CreateTaskListCompletedStatusInput) (*model.TaskListCompletedStatus, error)
	Update(ctx context.Context, input model.UpdateTaskListCompletedStatusInput) (*model.TaskListCompletedStatus, error)
}

type taskListCompletedStatusController struct {
	taskListCompletedStatusUsecase usecase.TaskListCompletedStatus
}

// NewTaskListCompletedStatusController generates taskListCompletedStatus controller.
func NewTaskListCompletedStatusController(u usecase.TaskListCompletedStatus) TaskListCompletedStatus {
	return &taskListCompletedStatusController{
		taskListCompletedStatusUsecase: u,
	}
}

func (c *taskListCompletedStatusController) Get(ctx context.Context, where *model.TaskListCompletedStatusWhereInput) (*model.TaskListCompletedStatus, error) {
	return c.taskListCompletedStatusUsecase.Get(ctx, where)
}

func (c *taskListCompletedStatusController) List(ctx context.Context) ([]*model.TaskListCompletedStatus, error) {
	return c.taskListCompletedStatusUsecase.List(ctx)
}

func (c *taskListCompletedStatusController) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskListCompletedStatusWhereInput, requestedFields []string) (*model.TaskListCompletedStatusConnection, error) {
	return c.taskListCompletedStatusUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (c *taskListCompletedStatusController) Create(ctx context.Context, input model.CreateTaskListCompletedStatusInput) (*model.TaskListCompletedStatus, error) {
	return c.taskListCompletedStatusUsecase.Create(ctx, input)
}

func (c *taskListCompletedStatusController) Update(ctx context.Context, input model.UpdateTaskListCompletedStatusInput) (*model.TaskListCompletedStatus, error) {
	return c.taskListCompletedStatusUsecase.Update(ctx, input)
}

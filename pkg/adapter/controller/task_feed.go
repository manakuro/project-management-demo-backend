package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TaskFeed is an interface of controller.
type TaskFeed interface {
	Get(ctx context.Context, where *model.TaskFeedWhereInput) (*model.TaskFeed, error)
	List(ctx context.Context, where *model.TaskFeedWhereInput) ([]*model.TaskFeed, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskFeedWhereInput, requestedFields []string) (*model.TaskFeedConnection, error)
	Create(ctx context.Context, input model.CreateTaskFeedInput) (*model.TaskFeed, error)
	Update(ctx context.Context, input model.UpdateTaskFeedInput) (*model.TaskFeed, error)
	Delete(ctx context.Context, input model.DeleteTaskFeedInput) (*model.TaskFeed, error)
}

type taskFeedController struct {
	taskFeedUsecase usecase.TaskFeed
}

// NewTaskFeedController generates controller.
func NewTaskFeedController(u usecase.TaskFeed) TaskFeed {
	return &taskFeedController{
		taskFeedUsecase: u,
	}
}

func (c *taskFeedController) Get(ctx context.Context, where *model.TaskFeedWhereInput) (*model.TaskFeed, error) {
	return c.taskFeedUsecase.Get(ctx, where)
}

func (c *taskFeedController) List(ctx context.Context, where *model.TaskFeedWhereInput) ([]*model.TaskFeed, error) {
	return c.taskFeedUsecase.List(ctx, where)
}

func (c *taskFeedController) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskFeedWhereInput, requestedFields []string) (*model.TaskFeedConnection, error) {
	return c.taskFeedUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (c *taskFeedController) Create(ctx context.Context, input model.CreateTaskFeedInput) (*model.TaskFeed, error) {
	return c.taskFeedUsecase.Create(ctx, input)
}

func (c *taskFeedController) Update(ctx context.Context, input model.UpdateTaskFeedInput) (*model.TaskFeed, error) {
	return c.taskFeedUsecase.Update(ctx, input)
}

func (c *taskFeedController) Delete(ctx context.Context, input model.DeleteTaskFeedInput) (*model.TaskFeed, error) {
	return c.taskFeedUsecase.Delete(ctx, input)
}

package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TaskSection is an interface of controller
type TaskSection interface {
	Get(ctx context.Context, where *model.TaskSectionWhereInput) (*model.TaskSection, error)
	List(ctx context.Context) ([]*model.TaskSection, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskSectionWhereInput, requestedFields []string) (*model.TaskSectionConnection, error)
	Create(ctx context.Context, input model.CreateTaskSectionInput) (*model.TaskSection, error)
	Update(ctx context.Context, input model.UpdateTaskSectionInput) (*model.TaskSection, error)
}

type taskSection struct {
	taskSectionUsecase usecase.TaskSection
}

// NewTaskSectionController generates taskSection controller
func NewTaskSectionController(pt usecase.TaskSection) TaskSection {
	return &taskSection{
		taskSectionUsecase: pt,
	}
}

func (t *taskSection) Get(ctx context.Context, where *model.TaskSectionWhereInput) (*model.TaskSection, error) {
	return t.taskSectionUsecase.Get(ctx, where)
}

func (t *taskSection) List(ctx context.Context) ([]*model.TaskSection, error) {
	return t.taskSectionUsecase.List(ctx)
}

func (t *taskSection) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskSectionWhereInput, requestedFields []string) (*model.TaskSectionConnection, error) {
	return t.taskSectionUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (t *taskSection) Create(ctx context.Context, input model.CreateTaskSectionInput) (*model.TaskSection, error) {
	return t.taskSectionUsecase.Create(ctx, input)
}

func (t *taskSection) Update(ctx context.Context, input model.UpdateTaskSectionInput) (*model.TaskSection, error) {
	return t.taskSectionUsecase.Update(ctx, input)
}

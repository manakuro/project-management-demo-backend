package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type taskSection struct {
	taskSectionRepository repository.TaskSection
}

// TaskSection is an interface of test user
type TaskSection interface {
	Get(ctx context.Context, where *model.TaskSectionWhereInput) (*model.TaskSection, error)
	List(ctx context.Context) ([]*model.TaskSection, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskSectionWhereInput, requestedFields []string) (*model.TaskSectionConnection, error)
	Create(ctx context.Context, input model.CreateTaskSectionInput) (*model.TaskSection, error)
	Update(ctx context.Context, input model.UpdateTaskSectionInput) (*model.TaskSection, error)
}

// NewTaskSectionUsecase generates test user repository
func NewTaskSectionUsecase(r repository.TaskSection) TaskSection {
	return &taskSection{taskSectionRepository: r}
}

func (t *taskSection) Get(ctx context.Context, where *model.TaskSectionWhereInput) (*model.TaskSection, error) {
	return t.taskSectionRepository.Get(ctx, where)
}

func (t *taskSection) List(ctx context.Context) ([]*model.TaskSection, error) {
	return t.taskSectionRepository.List(ctx)
}

func (t *taskSection) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskSectionWhereInput, requestedFields []string) (*model.TaskSectionConnection, error) {
	return t.taskSectionRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (t *taskSection) Create(ctx context.Context, input model.CreateTaskSectionInput) (*model.TaskSection, error) {
	return t.taskSectionRepository.Create(ctx, input)
}

func (t *taskSection) Update(ctx context.Context, input model.UpdateTaskSectionInput) (*model.TaskSection, error) {
	return t.taskSectionRepository.Update(ctx, input)
}

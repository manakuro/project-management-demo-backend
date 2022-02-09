package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type taskUsecase struct {
	taskRepository repository.Task
}

// Task is an interface of test user
type Task interface {
	Get(ctx context.Context, where *model.TaskWhereInput) (*model.Task, error)
	List(ctx context.Context) ([]*model.Task, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskWhereInput) (*model.TaskConnection, error)
	Create(ctx context.Context, input model.CreateTaskInput) (*model.Task, error)
	Update(ctx context.Context, input model.UpdateTaskInput) (*model.Task, error)
}

// NewTaskUsecase generates test user repository
func NewTaskUsecase(r repository.Task) Task {
	return &taskUsecase{taskRepository: r}
}

func (u *taskUsecase) Get(ctx context.Context, where *model.TaskWhereInput) (*model.Task, error) {
	return u.taskRepository.Get(ctx, where)
}

func (u *taskUsecase) List(ctx context.Context) ([]*model.Task, error) {
	return u.taskRepository.List(ctx)
}

func (u *taskUsecase) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskWhereInput) (*model.TaskConnection, error) {
	return u.taskRepository.ListWithPagination(ctx, after, first, before, last, where)
}

func (u *taskUsecase) Create(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	return u.taskRepository.Create(ctx, input)
}

func (u *taskUsecase) Update(ctx context.Context, input model.UpdateTaskInput) (*model.Task, error) {
	return u.taskRepository.Update(ctx, input)
}

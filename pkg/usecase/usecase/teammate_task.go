package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type teammateTaskUsecase struct {
	teammateTaskRepository repository.TeammateTask
}

// TeammateTask is an interface of usecase.
type TeammateTask interface {
	Get(ctx context.Context, where *model.TeammateTaskWhereInput) (*model.TeammateTask, error)
	List(ctx context.Context) ([]*model.TeammateTask, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskWhereInput, requestedFields []string) (*model.TeammateTaskConnection, error)
	Create(ctx context.Context, input model.CreateTeammateTaskInput) (*model.TeammateTask, error)
	Update(ctx context.Context, input model.UpdateTeammateTaskInput) (*model.TeammateTask, error)
	Delete(ctx context.Context, input model.DeleteTeammateTaskInput) (*model.TeammateTask, error)
}

// NewTeammateTaskUsecase generates a repository.
func NewTeammateTaskUsecase(r repository.TeammateTask) TeammateTask {
	return &teammateTaskUsecase{teammateTaskRepository: r}
}

func (u *teammateTaskUsecase) Get(ctx context.Context, where *model.TeammateTaskWhereInput) (*model.TeammateTask, error) {
	return u.teammateTaskRepository.Get(ctx, where)
}

func (u *teammateTaskUsecase) List(ctx context.Context) ([]*model.TeammateTask, error) {
	return u.teammateTaskRepository.List(ctx)
}

func (u *teammateTaskUsecase) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskWhereInput, requestedFields []string) (*model.TeammateTaskConnection, error) {
	return u.teammateTaskRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (u *teammateTaskUsecase) Create(ctx context.Context, input model.CreateTeammateTaskInput) (*model.TeammateTask, error) {
	return u.teammateTaskRepository.Create(ctx, input)
}

func (u *teammateTaskUsecase) Update(ctx context.Context, input model.UpdateTeammateTaskInput) (*model.TeammateTask, error) {
	return u.teammateTaskRepository.Update(ctx, input)
}

func (u *teammateTaskUsecase) Delete(ctx context.Context, input model.DeleteTeammateTaskInput) (*model.TeammateTask, error) {
	return u.teammateTaskRepository.Delete(ctx, input)
}

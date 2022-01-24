package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type projectIconUsecase struct {
	projectIconRepository repository.ProjectIcon
}

// ProjectIcon is an interface of test user
type ProjectIcon interface {
	Get(ctx context.Context, where *model.ProjectIconWhereInput) (*model.ProjectIcon, error)
	List(ctx context.Context) ([]*model.ProjectIcon, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectIconWhereInput, requestedFields []string) (*model.ProjectIconConnection, error)
	Create(ctx context.Context, input model.CreateProjectIconInput) (*model.ProjectIcon, error)
	Update(ctx context.Context, input model.UpdateProjectIconInput) (*model.ProjectIcon, error)
}

// NewProjectIconUsecase generates test user repository
func NewProjectIconUsecase(r repository.ProjectIcon) ProjectIcon {
	return &projectIconUsecase{projectIconRepository: r}
}

func (u *projectIconUsecase) Get(ctx context.Context, where *model.ProjectIconWhereInput) (*model.ProjectIcon, error) {
	return u.projectIconRepository.Get(ctx, where)
}

func (u *projectIconUsecase) List(ctx context.Context) ([]*model.ProjectIcon, error) {
	return u.projectIconRepository.List(ctx)
}

func (u *projectIconUsecase) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectIconWhereInput, requestedFields []string) (*model.ProjectIconConnection, error) {
	return u.projectIconRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (u *projectIconUsecase) Create(ctx context.Context, input model.CreateProjectIconInput) (*model.ProjectIcon, error) {
	return u.projectIconRepository.Create(ctx, input)
}

func (u *projectIconUsecase) Update(ctx context.Context, input model.UpdateProjectIconInput) (*model.ProjectIcon, error) {
	return u.projectIconRepository.Update(ctx, input)
}

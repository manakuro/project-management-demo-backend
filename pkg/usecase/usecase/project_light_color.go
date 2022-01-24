package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type projectLightColorUsecase struct {
	projectLightColorRepository repository.ProjectLightColor
}

// ProjectLightColor is an interface of test user
type ProjectLightColor interface {
	Get(ctx context.Context, where *model.ProjectLightColorWhereInput) (*model.ProjectLightColor, error)
	List(ctx context.Context) ([]*model.ProjectLightColor, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectLightColorWhereInput, requestedFields []string) (*model.ProjectLightColorConnection, error)
	Create(ctx context.Context, input model.CreateProjectLightColorInput) (*model.ProjectLightColor, error)
	Update(ctx context.Context, input model.UpdateProjectLightColorInput) (*model.ProjectLightColor, error)
}

// NewProjectLightColorUsecase generates test user repository
func NewProjectLightColorUsecase(r repository.ProjectLightColor) ProjectLightColor {
	return &projectLightColorUsecase{projectLightColorRepository: r}
}

func (u *projectLightColorUsecase) Get(ctx context.Context, where *model.ProjectLightColorWhereInput) (*model.ProjectLightColor, error) {
	return u.projectLightColorRepository.Get(ctx, where)
}

func (u *projectLightColorUsecase) List(ctx context.Context) ([]*model.ProjectLightColor, error) {
	return u.projectLightColorRepository.List(ctx)
}

func (u *projectLightColorUsecase) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectLightColorWhereInput, requestedFields []string) (*model.ProjectLightColorConnection, error) {
	return u.projectLightColorRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (u *projectLightColorUsecase) Create(ctx context.Context, input model.CreateProjectLightColorInput) (*model.ProjectLightColor, error) {
	return u.projectLightColorRepository.Create(ctx, input)
}

func (u *projectLightColorUsecase) Update(ctx context.Context, input model.UpdateProjectLightColorInput) (*model.ProjectLightColor, error) {
	return u.projectLightColorRepository.Update(ctx, input)
}

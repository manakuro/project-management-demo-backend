package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type projectLightColor struct {
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
	return &projectLightColor{projectLightColorRepository: r}
}

func (p *projectLightColor) Get(ctx context.Context, where *model.ProjectLightColorWhereInput) (*model.ProjectLightColor, error) {
	return p.projectLightColorRepository.Get(ctx, where)
}

func (p *projectLightColor) List(ctx context.Context) ([]*model.ProjectLightColor, error) {
	return p.projectLightColorRepository.List(ctx)
}

func (p *projectLightColor) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectLightColorWhereInput, requestedFields []string) (*model.ProjectLightColorConnection, error) {
	return p.projectLightColorRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (p *projectLightColor) Create(ctx context.Context, input model.CreateProjectLightColorInput) (*model.ProjectLightColor, error) {
	return p.projectLightColorRepository.Create(ctx, input)
}

func (p *projectLightColor) Update(ctx context.Context, input model.UpdateProjectLightColorInput) (*model.ProjectLightColor, error) {
	return p.projectLightColorRepository.Update(ctx, input)
}

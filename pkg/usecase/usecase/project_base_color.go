package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type projectBaseColor struct {
	projectBaseColorRepository repository.ProjectBaseColor
}

// ProjectBaseColor is an interface of test user
type ProjectBaseColor interface {
	Get(ctx context.Context, where *model.ProjectBaseColorWhereInput) (*model.ProjectBaseColor, error)
	List(ctx context.Context) ([]*model.ProjectBaseColor, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectBaseColorWhereInput, requestedFields []string) (*model.ProjectBaseColorConnection, error)
	Create(ctx context.Context, input model.CreateProjectBaseColorInput) (*model.ProjectBaseColor, error)
	Update(ctx context.Context, input model.UpdateProjectBaseColorInput) (*model.ProjectBaseColor, error)
}

// NewProjectBaseColorUsecase generates test user repository
func NewProjectBaseColorUsecase(r repository.ProjectBaseColor) ProjectBaseColor {
	return &projectBaseColor{projectBaseColorRepository: r}
}

func (p *projectBaseColor) Get(ctx context.Context, where *model.ProjectBaseColorWhereInput) (*model.ProjectBaseColor, error) {
	return p.projectBaseColorRepository.Get(ctx, where)
}

func (p *projectBaseColor) List(ctx context.Context) ([]*model.ProjectBaseColor, error) {
	return p.projectBaseColorRepository.List(ctx)
}

func (p *projectBaseColor) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectBaseColorWhereInput, requestedFields []string) (*model.ProjectBaseColorConnection, error) {
	return p.projectBaseColorRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (p *projectBaseColor) Create(ctx context.Context, input model.CreateProjectBaseColorInput) (*model.ProjectBaseColor, error) {
	return p.projectBaseColorRepository.Create(ctx, input)
}

func (p *projectBaseColor) Update(ctx context.Context, input model.UpdateProjectBaseColorInput) (*model.ProjectBaseColor, error) {
	return p.projectBaseColorRepository.Update(ctx, input)
}

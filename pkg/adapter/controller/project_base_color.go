package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// ProjectBaseColor is an interface of controller
type ProjectBaseColor interface {
	Get(ctx context.Context, where *model.ProjectBaseColorWhereInput) (*model.ProjectBaseColor, error)
	List(ctx context.Context) ([]*model.ProjectBaseColor, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectBaseColorWhereInput, requestedFields []string) (*model.ProjectBaseColorConnection, error)
	Create(ctx context.Context, input model.CreateProjectBaseColorInput) (*model.ProjectBaseColor, error)
	Update(ctx context.Context, input model.UpdateProjectBaseColorInput) (*model.ProjectBaseColor, error)
}

type projectBaseColor struct {
	projectBaseColorUsecase usecase.ProjectBaseColor
}

// NewProjectBaseColorController generates projectBaseColor controller
func NewProjectBaseColorController(p usecase.ProjectBaseColor) ProjectBaseColor {
	return &projectBaseColor{
		projectBaseColorUsecase: p,
	}
}

func (p *projectBaseColor) Get(ctx context.Context, where *model.ProjectBaseColorWhereInput) (*model.ProjectBaseColor, error) {
	return p.projectBaseColorUsecase.Get(ctx, where)
}

func (p *projectBaseColor) List(ctx context.Context) ([]*model.ProjectBaseColor, error) {
	return p.projectBaseColorUsecase.List(ctx)
}

func (p *projectBaseColor) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectBaseColorWhereInput, requestedFields []string) (*model.ProjectBaseColorConnection, error) {
	return p.projectBaseColorUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (p *projectBaseColor) Create(ctx context.Context, input model.CreateProjectBaseColorInput) (*model.ProjectBaseColor, error) {
	return p.projectBaseColorUsecase.Create(ctx, input)
}

func (p *projectBaseColor) Update(ctx context.Context, input model.UpdateProjectBaseColorInput) (*model.ProjectBaseColor, error) {
	return p.projectBaseColorUsecase.Update(ctx, input)
}

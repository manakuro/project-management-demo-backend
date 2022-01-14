package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// ProjectLightColor is an interface of controller
type ProjectLightColor interface {
	Get(ctx context.Context, where *model.ProjectLightColorWhereInput) (*model.ProjectLightColor, error)
	List(ctx context.Context) ([]*model.ProjectLightColor, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectLightColorWhereInput, requestedFields []string) (*model.ProjectLightColorConnection, error)
	Create(ctx context.Context, input model.CreateProjectLightColorInput) (*model.ProjectLightColor, error)
	Update(ctx context.Context, input model.UpdateProjectLightColorInput) (*model.ProjectLightColor, error)
}

type projectLightColor struct {
	projectLightColorUsecase usecase.ProjectLightColor
}

// NewProjectLightColorController generates projectLightColor controller
func NewProjectLightColorController(p usecase.ProjectLightColor) ProjectLightColor {
	return &projectLightColor{
		projectLightColorUsecase: p,
	}
}

func (p *projectLightColor) Get(ctx context.Context, where *model.ProjectLightColorWhereInput) (*model.ProjectLightColor, error) {
	return p.projectLightColorUsecase.Get(ctx, where)
}

func (p *projectLightColor) List(ctx context.Context) ([]*model.ProjectLightColor, error) {
	return p.projectLightColorUsecase.List(ctx)
}

func (p *projectLightColor) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectLightColorWhereInput, requestedFields []string) (*model.ProjectLightColorConnection, error) {
	return p.projectLightColorUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (p *projectLightColor) Create(ctx context.Context, input model.CreateProjectLightColorInput) (*model.ProjectLightColor, error) {
	return p.projectLightColorUsecase.Create(ctx, input)
}

func (p *projectLightColor) Update(ctx context.Context, input model.UpdateProjectLightColorInput) (*model.ProjectLightColor, error) {
	return p.projectLightColorUsecase.Update(ctx, input)
}

package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// ProjectIcon is an interface of controller
type ProjectIcon interface {
	Get(ctx context.Context, where *model.ProjectIconWhereInput) (*model.ProjectIcon, error)
	List(ctx context.Context) ([]*model.ProjectIcon, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectIconWhereInput, requestedFields []string) (*model.ProjectIconConnection, error)
	Create(ctx context.Context, input model.CreateProjectIconInput) (*model.ProjectIcon, error)
	Update(ctx context.Context, input model.UpdateProjectIconInput) (*model.ProjectIcon, error)
}

type projectIcon struct {
	projectIconUsecase usecase.ProjectIcon
}

// NewProjectIconController generates projectIcon controller
func NewProjectIconController(p usecase.ProjectIcon) ProjectIcon {
	return &projectIcon{
		projectIconUsecase: p,
	}
}

func (p *projectIcon) Get(ctx context.Context, where *model.ProjectIconWhereInput) (*model.ProjectIcon, error) {
	return p.projectIconUsecase.Get(ctx, where)
}

func (p *projectIcon) List(ctx context.Context) ([]*model.ProjectIcon, error) {
	return p.projectIconUsecase.List(ctx)
}

func (p *projectIcon) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectIconWhereInput, requestedFields []string) (*model.ProjectIconConnection, error) {
	return p.projectIconUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (p *projectIcon) Create(ctx context.Context, input model.CreateProjectIconInput) (*model.ProjectIcon, error) {
	return p.projectIconUsecase.Create(ctx, input)
}

func (p *projectIcon) Update(ctx context.Context, input model.UpdateProjectIconInput) (*model.ProjectIcon, error) {
	return p.projectIconUsecase.Update(ctx, input)
}

package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// Project is an interface of controller
type Project interface {
	Get(ctx context.Context, where *model.ProjectWhereInput) (*model.Project, error)
	List(ctx context.Context) ([]*model.Project, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectWhereInput) (*model.ProjectConnection, error)
	Create(ctx context.Context, input model.CreateProjectInput) (*model.Project, error)
	Update(ctx context.Context, input model.UpdateProjectInput) (*model.Project, error)
}

type project struct {
	projectUsecase usecase.Project
}

// NewProjectController generates project controller
func NewProjectController(tu usecase.Project) Project {
	return &project{
		projectUsecase: tu,
	}
}

func (t *project) Get(ctx context.Context, where *model.ProjectWhereInput) (*model.Project, error) {
	return t.projectUsecase.Get(ctx, where)
}

func (t *project) List(ctx context.Context) ([]*model.Project, error) {
	return t.projectUsecase.List(ctx)
}

func (t *project) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectWhereInput) (*model.ProjectConnection, error) {
	return t.projectUsecase.ListWithPagination(ctx, after, first, before, last, where)
}

func (t *project) Create(ctx context.Context, input model.CreateProjectInput) (*model.Project, error) {
	return t.projectUsecase.Create(ctx, input)
}

func (t *project) Update(ctx context.Context, input model.UpdateProjectInput) (*model.Project, error) {
	return t.projectUsecase.Update(ctx, input)
}

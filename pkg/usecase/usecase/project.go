package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type project struct {
	projectRepository repository.Project
}

// Project is an interface of test user
type Project interface {
	Get(ctx context.Context, where *model.ProjectWhereInput) (*model.Project, error)
	List(ctx context.Context) ([]*model.Project, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectWhereInput, requestedFields []string) (*model.ProjectConnection, error)
	Create(ctx context.Context, input model.CreateProjectInput) (*model.Project, error)
	Update(ctx context.Context, input model.UpdateProjectInput) (*model.Project, error)
}

// NewProjectUsecase generates test user repository
func NewProjectUsecase(r repository.Project) Project {
	return &project{projectRepository: r}
}

func (t *project) Get(ctx context.Context, where *model.ProjectWhereInput) (*model.Project, error) {
	return t.projectRepository.Get(ctx, where)
}

func (t *project) List(ctx context.Context) ([]*model.Project, error) {
	return t.projectRepository.List(ctx)
}

func (t *project) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectWhereInput, requestedFields []string) (*model.ProjectConnection, error) {
	return t.projectRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (t *project) Create(ctx context.Context, input model.CreateProjectInput) (*model.Project, error) {
	return t.projectRepository.Create(ctx, input)
}

func (t *project) Update(ctx context.Context, input model.UpdateProjectInput) (*model.Project, error) {
	return t.projectRepository.Update(ctx, input)
}

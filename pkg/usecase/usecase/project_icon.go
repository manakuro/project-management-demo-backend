package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type projectIcon struct {
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
	return &projectIcon{projectIconRepository: r}
}

func (p *projectIcon) Get(ctx context.Context, where *model.ProjectIconWhereInput) (*model.ProjectIcon, error) {
	return p.projectIconRepository.Get(ctx, where)
}

func (p *projectIcon) List(ctx context.Context) ([]*model.ProjectIcon, error) {
	return p.projectIconRepository.List(ctx)
}

func (p *projectIcon) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectIconWhereInput, requestedFields []string) (*model.ProjectIconConnection, error) {
	return p.projectIconRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (p *projectIcon) Create(ctx context.Context, input model.CreateProjectIconInput) (*model.ProjectIcon, error) {
	return p.projectIconRepository.Create(ctx, input)
}

func (p *projectIcon) Update(ctx context.Context, input model.UpdateProjectIconInput) (*model.ProjectIcon, error) {
	return p.projectIconRepository.Update(ctx, input)
}

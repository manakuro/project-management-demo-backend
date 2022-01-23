package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type projectTaskColumn struct {
	projectTaskColumnRepository repository.ProjectTaskColumn
}

// ProjectTaskColumn is an interface of test user
type ProjectTaskColumn interface {
	Get(ctx context.Context, where *model.ProjectTaskColumnWhereInput) (*model.ProjectTaskColumn, error)
	List(ctx context.Context) ([]*model.ProjectTaskColumn, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTaskColumnWhereInput, requestedFields []string) (*model.ProjectTaskColumnConnection, error)
	Create(ctx context.Context, input model.CreateProjectTaskColumnInput) (*model.ProjectTaskColumn, error)
	Update(ctx context.Context, input model.UpdateProjectTaskColumnInput) (*model.ProjectTaskColumn, error)
}

// NewProjectTaskColumnUsecase generates test user repository
func NewProjectTaskColumnUsecase(r repository.ProjectTaskColumn) ProjectTaskColumn {
	return &projectTaskColumn{projectTaskColumnRepository: r}
}

func (p *projectTaskColumn) Get(ctx context.Context, where *model.ProjectTaskColumnWhereInput) (*model.ProjectTaskColumn, error) {
	return p.projectTaskColumnRepository.Get(ctx, where)
}

func (p *projectTaskColumn) List(ctx context.Context) ([]*model.ProjectTaskColumn, error) {
	return p.projectTaskColumnRepository.List(ctx)
}

func (p *projectTaskColumn) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTaskColumnWhereInput, requestedFields []string) (*model.ProjectTaskColumnConnection, error) {
	return p.projectTaskColumnRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (p *projectTaskColumn) Create(ctx context.Context, input model.CreateProjectTaskColumnInput) (*model.ProjectTaskColumn, error) {
	return p.projectTaskColumnRepository.Create(ctx, input)
}

func (p *projectTaskColumn) Update(ctx context.Context, input model.UpdateProjectTaskColumnInput) (*model.ProjectTaskColumn, error) {
	return p.projectTaskColumnRepository.Update(ctx, input)
}

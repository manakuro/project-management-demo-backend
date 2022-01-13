package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type projectTeammate struct {
	projectTeammateRepository repository.ProjectTeammate
}

// ProjectTeammate is an interface of test user
type ProjectTeammate interface {
	Get(ctx context.Context, where *model.ProjectTeammateWhereInput) (*model.ProjectTeammate, error)
	List(ctx context.Context) ([]*model.ProjectTeammate, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTeammateWhereInput, requestedFields []string) (*model.ProjectTeammateConnection, error)
	Create(ctx context.Context, input model.CreateProjectTeammateInput) (*model.ProjectTeammate, error)
	Update(ctx context.Context, input model.UpdateProjectTeammateInput) (*model.ProjectTeammate, error)
}

// NewProjectTeammateUsecase generates test user repository
func NewProjectTeammateUsecase(r repository.ProjectTeammate) ProjectTeammate {
	return &projectTeammate{projectTeammateRepository: r}
}

func (p *projectTeammate) Get(ctx context.Context, where *model.ProjectTeammateWhereInput) (*model.ProjectTeammate, error) {
	return p.projectTeammateRepository.Get(ctx, where)
}

func (p *projectTeammate) List(ctx context.Context) ([]*model.ProjectTeammate, error) {
	return p.projectTeammateRepository.List(ctx)
}

func (p *projectTeammate) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTeammateWhereInput, requestedFields []string) (*model.ProjectTeammateConnection, error) {
	return p.projectTeammateRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (p *projectTeammate) Create(ctx context.Context, input model.CreateProjectTeammateInput) (*model.ProjectTeammate, error) {
	return p.projectTeammateRepository.Create(ctx, input)
}

func (p *projectTeammate) Update(ctx context.Context, input model.UpdateProjectTeammateInput) (*model.ProjectTeammate, error) {
	return p.projectTeammateRepository.Update(ctx, input)
}

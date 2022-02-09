package repository

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
)

// ProjectTeammate is interface of repository
type ProjectTeammate interface {
	Get(ctx context.Context, where *model.ProjectTeammateWhereInput) (*model.ProjectTeammate, error)
	List(ctx context.Context) ([]*model.ProjectTeammate, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTeammateWhereInput) (*model.ProjectTeammateConnection, error)
	Create(ctx context.Context, input model.CreateProjectTeammateInput) (*model.ProjectTeammate, error)
	Update(ctx context.Context, input model.UpdateProjectTeammateInput) (*model.ProjectTeammate, error)
}

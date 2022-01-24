package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type workspaceTeammateRepository struct {
	client *ent.Client
}

// NewWorkspaceTeammateRepository generates workspaceTeammate repository
func NewWorkspaceTeammateRepository(client *ent.Client) ur.WorkspaceTeammate {
	return &workspaceTeammateRepository{client: client}
}

func (r *workspaceTeammateRepository) Get(ctx context.Context, where *model.WorkspaceTeammateWhereInput) (*model.WorkspaceTeammate, error) {
	q := r.client.WorkspaceTeammate.Query()

	q, err := where.Filter(q)
	if err != nil {
		return nil, model.NewInvalidParamError(nil)
	}

	result, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, nil)
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (r *workspaceTeammateRepository) List(ctx context.Context) ([]*model.WorkspaceTeammate, error) {
	result, err := r.client.
		WorkspaceTeammate.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (r *workspaceTeammateRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceTeammateWhereInput, requestedFields []string) (*model.WorkspaceTeammateConnection, error) {
	q := r.client.WorkspaceTeammate.Query()

	if collection.Contains(requestedFields, "edges.node.workspace") {
		q.WithWorkspace()
	}

	if collection.Contains(requestedFields, "edges.node.teammate") {
		q.WithTeammate()
	}

	result, err := q.Paginate(ctx, after, first, before, last, ent.WithWorkspaceTeammateFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return result, nil
}

func (r *workspaceTeammateRepository) Create(ctx context.Context, input model.CreateWorkspaceTeammateInput) (*model.WorkspaceTeammate, error) {
	result, err := r.client.
		WorkspaceTeammate.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (r *workspaceTeammateRepository) Update(ctx context.Context, input model.UpdateWorkspaceTeammateInput) (*model.WorkspaceTeammate, error) {
	result, err := r.client.
		WorkspaceTeammate.UpdateOneID(input.ID).
		SetInput(input).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}

		return nil, model.NewDBError(err)
	}

	return result, nil
}

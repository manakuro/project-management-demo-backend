package workspaceteammaterepository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *workspaceTeammateRepository) Get(ctx context.Context, where *model.WorkspaceTeammateWhereInput) (*model.WorkspaceTeammate, error) {
	q := r.client.WorkspaceTeammate.Query()

	q, err := where.Filter(q)
	if err != nil {
		return nil, model.NewInvalidParamError(nil)
	}

	res, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, nil)
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *workspaceTeammateRepository) List(ctx context.Context) ([]*model.WorkspaceTeammate, error) {
	res, err := r.client.
		WorkspaceTeammate.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *workspaceTeammateRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceTeammateWhereInput) (*model.WorkspaceTeammateConnection, error) {
	q := r.client.WorkspaceTeammate.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithWorkspaceTeammateFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

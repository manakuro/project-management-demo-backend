package projectteammaterepository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *projectTeammateRepository) Get(ctx context.Context, where *model.ProjectTeammateWhereInput) (*model.ProjectTeammate, error) {
	q := r.client.ProjectTeammate.Query()

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

func (r *projectTeammateRepository) List(ctx context.Context) ([]*model.ProjectTeammate, error) {
	res, err := r.client.
		ProjectTeammate.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectTeammateRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTeammateWhereInput) (*model.ProjectTeammateConnection, error) {
	q := r.client.ProjectTeammate.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithProjectTeammateFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

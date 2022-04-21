package archivedworkspaceacivitytaskrepository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *archivedWorkspaceActivityTaskRepository) Get(ctx context.Context, where *model.ArchivedWorkspaceActivityTaskWhereInput) (*model.ArchivedWorkspaceActivityTask, error) {
	q := r.client.ArchivedWorkspaceActivityTask.Query()

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

func (r *archivedWorkspaceActivityTaskRepository) List(ctx context.Context) ([]*model.ArchivedWorkspaceActivityTask, error) {
	res, err := r.client.ArchivedWorkspaceActivityTask.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *archivedWorkspaceActivityTaskRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ArchivedWorkspaceActivityTaskWhereInput) (*model.ArchivedWorkspaceActivityTaskConnection, error) {
	q := r.client.ArchivedWorkspaceActivityTask.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithArchivedWorkspaceActivityTaskFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

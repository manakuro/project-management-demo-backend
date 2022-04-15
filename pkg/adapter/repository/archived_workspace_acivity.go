package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type archivedWorkspaceActivityRepository struct {
	client *ent.Client
}

// NewArchivedWorkspaceActivityRepository generates archivedWorkspaceActivity repository
func NewArchivedWorkspaceActivityRepository(client *ent.Client) ur.ArchivedWorkspaceActivity {
	return &archivedWorkspaceActivityRepository{client: client}
}

func (r *archivedWorkspaceActivityRepository) Get(ctx context.Context, where *model.ArchivedWorkspaceActivityWhereInput) (*model.ArchivedWorkspaceActivity, error) {
	q := r.client.ArchivedWorkspaceActivity.Query()

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

func (r *archivedWorkspaceActivityRepository) List(ctx context.Context) ([]*model.ArchivedWorkspaceActivity, error) {
	res, err := r.client.ArchivedWorkspaceActivity.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *archivedWorkspaceActivityRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ArchivedWorkspaceActivityWhereInput) (*model.ArchivedWorkspaceActivityConnection, error) {
	q := r.client.ArchivedWorkspaceActivity.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithArchivedWorkspaceActivityFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *archivedWorkspaceActivityRepository) Create(ctx context.Context, input model.CreateArchivedWorkspaceActivityInput) (*model.ArchivedWorkspaceActivity, error) {
	res, err := r.client.
		ArchivedWorkspaceActivity.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *archivedWorkspaceActivityRepository) Update(ctx context.Context, input model.UpdateArchivedWorkspaceActivityInput) (*model.ArchivedWorkspaceActivity, error) {
	res, err := r.client.
		ArchivedWorkspaceActivity.UpdateOneID(input.ID).
		SetInput(input).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}

		return nil, model.NewDBError(err)
	}

	return res, nil
}

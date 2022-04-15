package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type archivedTaskActivityRepository struct {
	client *ent.Client
}

// NewArchivedTaskActivityRepository generates archivedTaskActivity repository
func NewArchivedTaskActivityRepository(client *ent.Client) ur.ArchivedTaskActivity {
	return &archivedTaskActivityRepository{client: client}
}

func (r *archivedTaskActivityRepository) Get(ctx context.Context, where *model.ArchivedTaskActivityWhereInput) (*model.ArchivedTaskActivity, error) {
	q := r.client.ArchivedTaskActivity.Query()

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

func (r *archivedTaskActivityRepository) List(ctx context.Context) ([]*model.ArchivedTaskActivity, error) {
	res, err := r.client.ArchivedTaskActivity.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *archivedTaskActivityRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ArchivedTaskActivityWhereInput) (*model.ArchivedTaskActivityConnection, error) {
	q := r.client.ArchivedTaskActivity.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithArchivedTaskActivityFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *archivedTaskActivityRepository) Create(ctx context.Context, input model.CreateArchivedTaskActivityInput) (*model.ArchivedTaskActivity, error) {
	res, err := r.client.
		ArchivedTaskActivity.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *archivedTaskActivityRepository) Update(ctx context.Context, input model.UpdateArchivedTaskActivityInput) (*model.ArchivedTaskActivity, error) {
	res, err := r.client.
		ArchivedTaskActivity.UpdateOneID(input.ID).
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

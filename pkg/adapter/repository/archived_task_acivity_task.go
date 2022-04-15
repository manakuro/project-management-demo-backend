package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type archivedArchivedTaskActivityTaskRepository struct {
	client *ent.Client
}

// NewArchivedTaskActivityTaskRepository generates archivedArchivedTaskActivityTask repository
func NewArchivedTaskActivityTaskRepository(client *ent.Client) ur.ArchivedTaskActivityTask {
	return &archivedArchivedTaskActivityTaskRepository{client: client}
}

func (r *archivedArchivedTaskActivityTaskRepository) Get(ctx context.Context, where *model.ArchivedTaskActivityTaskWhereInput) (*model.ArchivedTaskActivityTask, error) {
	q := r.client.ArchivedTaskActivityTask.Query()

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

func (r *archivedArchivedTaskActivityTaskRepository) List(ctx context.Context) ([]*model.ArchivedTaskActivityTask, error) {
	res, err := r.client.ArchivedTaskActivityTask.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *archivedArchivedTaskActivityTaskRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ArchivedTaskActivityTaskWhereInput) (*model.ArchivedTaskActivityTaskConnection, error) {
	q := r.client.ArchivedTaskActivityTask.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithArchivedTaskActivityTaskFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *archivedArchivedTaskActivityTaskRepository) Create(ctx context.Context, input model.CreateArchivedTaskActivityTaskInput) (*model.ArchivedTaskActivityTask, error) {
	res, err := r.client.
		ArchivedTaskActivityTask.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *archivedArchivedTaskActivityTaskRepository) Update(ctx context.Context, input model.UpdateArchivedTaskActivityTaskInput) (*model.ArchivedTaskActivityTask, error) {
	res, err := r.client.
		ArchivedTaskActivityTask.UpdateOneID(input.ID).
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

package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskColumnRepository struct {
	client *ent.Client
}

// NewTaskColumnRepository generates taskColumn repository
func NewTaskColumnRepository(client *ent.Client) ur.TaskColumn {
	return &taskColumnRepository{client: client}
}

func (r *taskColumnRepository) Get(ctx context.Context, where *model.TaskColumnWhereInput) (*model.TaskColumn, error) {
	q := r.client.TaskColumn.Query()

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

func (r *taskColumnRepository) List(ctx context.Context) ([]*model.TaskColumn, error) {
	res, err := r.client.TaskColumn.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskColumnRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskColumnWhereInput) (*model.TaskColumnConnection, error) {
	q := r.client.TaskColumn.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskColumnFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskColumnRepository) Create(ctx context.Context, input model.CreateTaskColumnInput) (*model.TaskColumn, error) {
	res, err := r.client.
		TaskColumn.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskColumnRepository) Update(ctx context.Context, input model.UpdateTaskColumnInput) (*model.TaskColumn, error) {
	res, err := r.client.
		TaskColumn.UpdateOneID(input.ID).
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

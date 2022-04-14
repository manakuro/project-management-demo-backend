package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskActivityRepository struct {
	client *ent.Client
}

// NewTaskActivityRepository generates taskActivity repository
func NewTaskActivityRepository(client *ent.Client) ur.TaskActivity {
	return &taskActivityRepository{client: client}
}

func (r *taskActivityRepository) Get(ctx context.Context, where *model.TaskActivityWhereInput) (*model.TaskActivity, error) {
	q := r.client.TaskActivity.Query()

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

func (r *taskActivityRepository) List(ctx context.Context) ([]*model.TaskActivity, error) {
	res, err := r.client.TaskActivity.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskActivityRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskActivityWhereInput) (*model.TaskActivityConnection, error) {
	q := r.client.TaskActivity.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskActivityFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskActivityRepository) Create(ctx context.Context, input model.CreateTaskActivityInput) (*model.TaskActivity, error) {
	res, err := r.client.
		TaskActivity.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskActivityRepository) Update(ctx context.Context, input model.UpdateTaskActivityInput) (*model.TaskActivity, error) {
	res, err := r.client.
		TaskActivity.UpdateOneID(input.ID).
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

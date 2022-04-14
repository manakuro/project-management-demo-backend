package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskActivityTaskRepository struct {
	client *ent.Client
}

// NewTaskActivityTaskRepository generates taskActivityTask repository
func NewTaskActivityTaskRepository(client *ent.Client) ur.TaskActivityTask {
	return &taskActivityTaskRepository{client: client}
}

func (r *taskActivityTaskRepository) Get(ctx context.Context, where *model.TaskActivityTaskWhereInput) (*model.TaskActivityTask, error) {
	q := r.client.TaskActivityTask.Query()

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

func (r *taskActivityTaskRepository) List(ctx context.Context) ([]*model.TaskActivityTask, error) {
	res, err := r.client.TaskActivityTask.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskActivityTaskRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskActivityTaskWhereInput) (*model.TaskActivityTaskConnection, error) {
	q := r.client.TaskActivityTask.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskActivityTaskFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskActivityTaskRepository) Create(ctx context.Context, input model.CreateTaskActivityTaskInput) (*model.TaskActivityTask, error) {
	res, err := r.client.
		TaskActivityTask.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskActivityTaskRepository) Update(ctx context.Context, input model.UpdateTaskActivityTaskInput) (*model.TaskActivityTask, error) {
	res, err := r.client.
		TaskActivityTask.UpdateOneID(input.ID).
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

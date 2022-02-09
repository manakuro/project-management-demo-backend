package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskPriorityRepository struct {
	client *ent.Client
}

// NewTaskPriorityRepository generates taskPriority repository
func NewTaskPriorityRepository(client *ent.Client) ur.TaskPriority {
	return &taskPriorityRepository{client: client}
}

func (r *taskPriorityRepository) Get(ctx context.Context, where *model.TaskPriorityWhereInput) (*model.TaskPriority, error) {
	q := r.client.TaskPriority.Query()

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

func (r *taskPriorityRepository) List(ctx context.Context) ([]*model.TaskPriority, error) {
	res, err := r.client.TaskPriority.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskPriorityRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskPriorityWhereInput) (*model.TaskPriorityConnection, error) {
	q := r.client.TaskPriority.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskPriorityFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskPriorityRepository) Create(ctx context.Context, input model.CreateTaskPriorityInput) (*model.TaskPriority, error) {
	res, err := r.client.
		TaskPriority.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskPriorityRepository) Update(ctx context.Context, input model.UpdateTaskPriorityInput) (*model.TaskPriority, error) {
	res, err := r.client.
		TaskPriority.UpdateOneID(input.ID).
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

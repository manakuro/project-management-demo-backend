package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskListSortStatusRepository struct {
	client *ent.Client
}

// NewTaskListSortStatusRepository generates taskListSortStatus repository
func NewTaskListSortStatusRepository(client *ent.Client) ur.TaskListSortStatus {
	return &taskListSortStatusRepository{client: client}
}

func (r *taskListSortStatusRepository) Get(ctx context.Context, where *model.TaskListSortStatusWhereInput) (*model.TaskListSortStatus, error) {
	q := r.client.TaskListSortStatus.Query()

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

func (r *taskListSortStatusRepository) List(ctx context.Context) ([]*model.TaskListSortStatus, error) {
	res, err := r.client.TaskListSortStatus.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskListSortStatusRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskListSortStatusWhereInput, requestedFields []string) (*model.TaskListSortStatusConnection, error) {
	q := r.client.TaskListSortStatus.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskListSortStatusFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskListSortStatusRepository) Create(ctx context.Context, input model.CreateTaskListSortStatusInput) (*model.TaskListSortStatus, error) {
	res, err := r.client.
		TaskListSortStatus.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskListSortStatusRepository) Update(ctx context.Context, input model.UpdateTaskListSortStatusInput) (*model.TaskListSortStatus, error) {
	res, err := r.client.
		TaskListSortStatus.UpdateOneID(input.ID).
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

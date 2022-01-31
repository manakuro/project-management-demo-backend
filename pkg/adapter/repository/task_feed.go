package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/taskfeed"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskFeedRepository struct {
	client *ent.Client
}

// NewTaskFeedRepository generates taskFeed repository
func NewTaskFeedRepository(client *ent.Client) ur.TaskFeed {
	return &taskFeedRepository{client: client}
}

func (r *taskFeedRepository) Get(ctx context.Context, where *model.TaskFeedWhereInput) (*model.TaskFeed, error) {
	q := r.client.TaskFeed.Query()

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

func (r *taskFeedRepository) List(ctx context.Context, where *model.TaskFeedWhereInput) ([]*model.TaskFeed, error) {
	q := r.client.TaskFeed.Query()

	q, err := where.Filter(q)
	if err != nil {
		return nil, model.NewInvalidParamError(nil)
	}

	res, err := q.All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskFeedRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskFeedWhereInput, requestedFields []string) (*model.TaskFeedConnection, error) {
	q := r.client.TaskFeed.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskFeedFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskFeedRepository) Create(ctx context.Context, input model.CreateTaskFeedInput) (*model.TaskFeed, error) {
	res, err := r.client.
		TaskFeed.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskFeedRepository) Update(ctx context.Context, input model.UpdateTaskFeedInput) (*model.TaskFeed, error) {
	res, err := r.client.
		TaskFeed.UpdateOneID(input.ID).
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

func (r *taskFeedRepository) Delete(ctx context.Context, input model.DeleteTaskFeedInput) (*model.TaskFeed, error) {
	deleted, err := r.client.TaskFeed.Query().Where(taskfeed.IDEQ(input.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = r.client.TaskFeed.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}

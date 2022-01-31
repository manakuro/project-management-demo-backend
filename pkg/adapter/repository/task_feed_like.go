package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/taskfeedlike"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskFeedLikeRepository struct {
	client *ent.Client
}

// NewTaskFeedLikeRepository generates taskFeedLike repository
func NewTaskFeedLikeRepository(client *ent.Client) ur.TaskFeedLike {
	return &taskFeedLikeRepository{client: client}
}

func (r *taskFeedLikeRepository) Get(ctx context.Context, where *model.TaskFeedLikeWhereInput) (*model.TaskFeedLike, error) {
	q := r.client.TaskFeedLike.Query()

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

func (r *taskFeedLikeRepository) List(ctx context.Context, where *model.TaskFeedLikeWhereInput) ([]*model.TaskFeedLike, error) {
	q := r.client.TaskFeedLike.Query()

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

func (r *taskFeedLikeRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskFeedLikeWhereInput, requestedFields []string) (*model.TaskFeedLikeConnection, error) {
	q := r.client.TaskFeedLike.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskFeedLikeFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskFeedLikeRepository) Create(ctx context.Context, input model.CreateTaskFeedLikeInput) (*model.TaskFeedLike, error) {
	res, err := r.client.
		TaskFeedLike.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskFeedLikeRepository) Update(ctx context.Context, input model.UpdateTaskFeedLikeInput) (*model.TaskFeedLike, error) {
	res, err := r.client.
		TaskFeedLike.UpdateOneID(input.ID).
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

func (r *taskFeedLikeRepository) Delete(ctx context.Context, input model.DeleteTaskFeedLikeInput) (*model.TaskFeedLike, error) {
	deleted, err := r.client.TaskFeedLike.Query().Where(taskfeedlike.IDEQ(input.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = r.client.TaskFeedLike.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}

package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/tasklike"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskLikeRepository struct {
	client *ent.Client
}

// NewTaskLikeRepository generates taskLike repository
func NewTaskLikeRepository(client *ent.Client) ur.TaskLike {
	return &taskLikeRepository{client: client}
}

func (r *taskLikeRepository) Get(ctx context.Context, where *model.TaskLikeWhereInput) (*model.TaskLike, error) {
	q := r.client.TaskLike.Query()

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

func (r *taskLikeRepository) List(ctx context.Context, where *model.TaskLikeWhereInput) ([]*model.TaskLike, error) {
	q := r.client.TaskLike.Query()

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

func (r *taskLikeRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskLikeWhereInput) (*model.TaskLikeConnection, error) {
	q := r.client.TaskLike.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskLikeFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskLikeRepository) Create(ctx context.Context, input model.CreateTaskLikeInput) (*model.TaskLike, error) {
	res, err := r.client.
		TaskLike.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskLikeRepository) Update(ctx context.Context, input model.UpdateTaskLikeInput) (*model.TaskLike, error) {
	res, err := r.client.
		TaskLike.UpdateOneID(input.ID).
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

func (r *taskLikeRepository) Delete(ctx context.Context, input model.DeleteTaskLikeInput) (*model.TaskLike, error) {
	deleted, err := r.client.TaskLike.Query().Where(tasklike.IDEQ(input.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = r.client.TaskLike.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}

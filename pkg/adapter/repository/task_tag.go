package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/tasktag"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskTagRepository struct {
	client *ent.Client
}

// NewTaskTagRepository generates taskTag repository
func NewTaskTagRepository(client *ent.Client) ur.TaskTag {
	return &taskTagRepository{client: client}
}

func (r *taskTagRepository) Get(ctx context.Context, where *model.TaskTagWhereInput) (*model.TaskTag, error) {
	q := r.client.TaskTag.Query()

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

func (r *taskTagRepository) List(ctx context.Context, where *model.TaskTagWhereInput) ([]*model.TaskTag, error) {
	q := r.client.TaskTag.Query()

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

func (r *taskTagRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskTagWhereInput) (*model.TaskTagConnection, error) {
	q := r.client.TaskTag.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskTagFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskTagRepository) Create(ctx context.Context, input model.CreateTaskTagInput) (*model.TaskTag, error) {
	client := WithTransactionalMutation(ctx)

	res, err := client.
		TaskTag.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	taskTag, err := client.TaskTag.Query().
		Where(tasktag.ID(res.ID)).
		WithTag(func(tq *ent.TagQuery) {
			WithTag(tq)
		}).Only(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return taskTag.Unwrap(), nil
}

func (r *taskTagRepository) Update(ctx context.Context, input model.UpdateTaskTagInput) (*model.TaskTag, error) {
	res, err := r.client.
		TaskTag.UpdateOneID(input.ID).
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

func (r *taskTagRepository) Delete(ctx context.Context, input model.DeleteTaskTagInput) (*model.TaskTag, error) {
	client := WithTransactionalMutation(ctx)

	deleted, err := client.TaskTag.Query().WithTag(func(tq *ent.TagQuery) {
		WithTag(tq)
	}).Where(tasktag.IDEQ(input.ID)).Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = client.TaskTag.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}

// WithTaskTag eager-loads associations with task tag entity.
func WithTaskTag(query *ent.TaskTagQuery) {
	query.WithTag(func(tq *ent.TagQuery) {
		WithTag(tq)
	})
}

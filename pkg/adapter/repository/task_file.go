package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskFileRepository struct {
	client *ent.Client
}

// NewTaskFileRepository generates taskFile repository
func NewTaskFileRepository(client *ent.Client) ur.TaskFile {
	return &taskFileRepository{client: client}
}

func (r *taskFileRepository) Get(ctx context.Context, where *model.TaskFileWhereInput) (*model.TaskFile, error) {
	q := r.client.TaskFile.Query()

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

func (r *taskFileRepository) List(ctx context.Context) ([]*model.TaskFile, error) {
	res, err := r.client.TaskFile.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskFileRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskFileWhereInput) (*model.TaskFileConnection, error) {
	q := r.client.TaskFile.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskFileFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskFileRepository) Create(ctx context.Context, input model.CreateTaskFileInput) (*model.TaskFile, error) {
	res, err := r.client.
		TaskFile.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskFileRepository) Update(ctx context.Context, input model.UpdateTaskFileInput) (*model.TaskFile, error) {
	res, err := r.client.
		TaskFile.UpdateOneID(input.ID).
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

// WithTaskFiles eager-loads association with task files collaborators.
func WithTaskFiles(query *ent.TaskFileQuery) {
	query.WithFileType()
}

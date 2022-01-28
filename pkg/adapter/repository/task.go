package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type taskRepository struct {
	client *ent.Client
}

// NewTaskRepository generates task repository
func NewTaskRepository(client *ent.Client) ur.Task {
	return &taskRepository{client: client}
}

func (r *taskRepository) Get(ctx context.Context, where *model.TaskWhereInput) (*model.Task, error) {
	q := r.client.Task.Query()

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

func (r *taskRepository) List(ctx context.Context) ([]*model.Task, error) {
	res, err := r.client.Task.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskWhereInput, requestedFields []string) (*model.TaskConnection, error) {
	q := r.client.Task.Query()

	if collection.Contains(requestedFields, "edges.node.taskPriority") {
		q.WithTaskPriority()
	}

	if collection.Contains(requestedFields, "edges.node.subTasks") {
		q.WithSubTasks()
	}

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskRepository) Create(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	res, err := r.client.
		Task.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskRepository) Update(ctx context.Context, input model.UpdateTaskInput) (*model.Task, error) {
	res, err := r.client.
		Task.UpdateOneID(input.ID).
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

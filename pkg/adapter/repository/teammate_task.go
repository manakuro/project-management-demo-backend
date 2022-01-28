package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/teammatetask"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type teammateTaskRepository struct {
	client *ent.Client
}

// NewTeammateTaskRepository generates teammateTask repository
func NewTeammateTaskRepository(client *ent.Client) ur.TeammateTask {
	return &teammateTaskRepository{client: client}
}

func (r *teammateTaskRepository) Get(ctx context.Context, where *model.TeammateTaskWhereInput) (*model.TeammateTask, error) {
	q := r.client.TeammateTask.Query()

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

func (r *teammateTaskRepository) List(ctx context.Context) ([]*model.TeammateTask, error) {
	res, err := r.client.TeammateTask.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTaskRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateTaskWhereInput, requestedFields []string) (*model.TeammateTaskConnection, error) {
	q := r.client.TeammateTask.Query()

	if collection.Contains(requestedFields, "edges.node.task") {
		q.WithTask(func(query *ent.TaskQuery) {
			query.WithSubTasks()
		})
	}

	if collection.Contains(requestedFields, "edges.node.teammateTaskSection") {
		q.WithTeammateTaskSection()
	}

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTeammateTaskFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *teammateTaskRepository) Create(ctx context.Context, input model.CreateTeammateTaskInput) (*model.TeammateTask, error) {
	res, err := r.client.
		TeammateTask.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *teammateTaskRepository) Update(ctx context.Context, input model.UpdateTeammateTaskInput) (*model.TeammateTask, error) {
	res, err := r.client.
		TeammateTask.UpdateOneID(input.ID).
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

func (r *teammateTaskRepository) Delete(ctx context.Context, input model.DeleteTeammateTaskInput) (*model.TeammateTask, error) {
	deleted, err := r.client.TeammateTask.Query().Where(teammatetask.IDEQ(input.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = r.client.TeammateTask.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}

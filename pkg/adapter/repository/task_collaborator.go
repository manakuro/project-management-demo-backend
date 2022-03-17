package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/taskcollaborator"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type taskCollaboratorRepository struct {
	client *ent.Client
}

// NewTaskCollaboratorRepository generates taskCollaborator repository
func NewTaskCollaboratorRepository(client *ent.Client) ur.TaskCollaborator {
	return &taskCollaboratorRepository{client: client}
}

func (r *taskCollaboratorRepository) Get(ctx context.Context, where *model.TaskCollaboratorWhereInput) (*model.TaskCollaborator, error) {
	q := r.client.TaskCollaborator.Query()

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

func (r *taskCollaboratorRepository) List(ctx context.Context, where *model.TaskCollaboratorWhereInput) ([]*model.TaskCollaborator, error) {
	q := r.client.TaskCollaborator.Query()

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

func (r *taskCollaboratorRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TaskCollaboratorWhereInput) (*model.TaskCollaboratorConnection, error) {
	q := r.client.TaskCollaborator.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTaskCollaboratorFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *taskCollaboratorRepository) Create(ctx context.Context, input model.CreateTaskCollaboratorInput) (*model.TaskCollaborator, error) {
	client := WithTransactionalMutation(ctx)

	res, err := client.
		TaskCollaborator.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *taskCollaboratorRepository) Update(ctx context.Context, input model.UpdateTaskCollaboratorInput) (*model.TaskCollaborator, error) {
	res, err := r.client.
		TaskCollaborator.UpdateOneID(input.ID).
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

func (r *taskCollaboratorRepository) Delete(ctx context.Context, input model.DeleteTaskCollaboratorInput) (*model.TaskCollaborator, error) {
	deleted, err := r.client.TaskCollaborator.Query().Where(taskcollaborator.IDEQ(input.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}
		return nil, model.NewDBError(err)
	}

	err = r.client.TaskCollaborator.DeleteOneID(input.ID).Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}

package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type workspaceActivityTaskRepository struct {
	client *ent.Client
}

// NewWorkspaceActivityTaskRepository generates workspaceActivityTask repository
func NewWorkspaceActivityTaskRepository(client *ent.Client) ur.WorkspaceActivityTask {
	return &workspaceActivityTaskRepository{client: client}
}

func (r *workspaceActivityTaskRepository) Get(ctx context.Context, where *model.WorkspaceActivityTaskWhereInput) (*model.WorkspaceActivityTask, error) {
	q := r.client.WorkspaceActivityTask.Query()

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

func (r *workspaceActivityTaskRepository) List(ctx context.Context) ([]*model.WorkspaceActivityTask, error) {
	res, err := r.client.WorkspaceActivityTask.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *workspaceActivityTaskRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceActivityTaskWhereInput) (*model.WorkspaceActivityTaskConnection, error) {
	q := r.client.WorkspaceActivityTask.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithWorkspaceActivityTaskFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *workspaceActivityTaskRepository) Create(ctx context.Context, input model.CreateWorkspaceActivityTaskInput) (*model.WorkspaceActivityTask, error) {
	res, err := r.client.
		WorkspaceActivityTask.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *workspaceActivityTaskRepository) Update(ctx context.Context, input model.UpdateWorkspaceActivityTaskInput) (*model.WorkspaceActivityTask, error) {
	res, err := r.client.
		WorkspaceActivityTask.UpdateOneID(input.ID).
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

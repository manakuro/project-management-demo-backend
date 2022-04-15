package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type workspaceActivityRepository struct {
	client *ent.Client
}

// NewWorkspaceActivityRepository generates workspaceActivity repository
func NewWorkspaceActivityRepository(client *ent.Client) ur.WorkspaceActivity {
	return &workspaceActivityRepository{client: client}
}

func (r *workspaceActivityRepository) Get(ctx context.Context, where *model.WorkspaceActivityWhereInput) (*model.WorkspaceActivity, error) {
	q := r.client.WorkspaceActivity.Query()

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

func (r *workspaceActivityRepository) List(ctx context.Context) ([]*model.WorkspaceActivity, error) {
	res, err := r.client.WorkspaceActivity.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *workspaceActivityRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.WorkspaceActivityWhereInput) (*model.WorkspaceActivityConnection, error) {
	q := r.client.WorkspaceActivity.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithWorkspaceActivityFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *workspaceActivityRepository) Create(ctx context.Context, input model.CreateWorkspaceActivityInput) (*model.WorkspaceActivity, error) {
	res, err := r.client.
		WorkspaceActivity.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *workspaceActivityRepository) Update(ctx context.Context, input model.UpdateWorkspaceActivityInput) (*model.WorkspaceActivity, error) {
	res, err := r.client.
		WorkspaceActivity.UpdateOneID(input.ID).
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

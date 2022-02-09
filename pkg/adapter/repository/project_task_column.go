package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type projectTaskColumnRepository struct {
	client *ent.Client
}

// NewProjectTaskColumnRepository generates projectTaskColumn repository
func NewProjectTaskColumnRepository(client *ent.Client) ur.ProjectTaskColumn {
	return &projectTaskColumnRepository{client: client}
}

func (r *projectTaskColumnRepository) Get(ctx context.Context, where *model.ProjectTaskColumnWhereInput) (*model.ProjectTaskColumn, error) {
	q := r.client.ProjectTaskColumn.Query()

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

func (r *projectTaskColumnRepository) List(ctx context.Context) ([]*model.ProjectTaskColumn, error) {
	res, err := r.client.ProjectTaskColumn.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectTaskColumnRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTaskColumnWhereInput, requestedFields []string) (*model.ProjectTaskColumnConnection, error) {
	q := r.client.ProjectTaskColumn.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithProjectTaskColumnFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *projectTaskColumnRepository) Create(ctx context.Context, input model.CreateProjectTaskColumnInput) (*model.ProjectTaskColumn, error) {
	res, err := r.client.
		ProjectTaskColumn.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectTaskColumnRepository) Update(ctx context.Context, input model.UpdateProjectTaskColumnInput) (*model.ProjectTaskColumn, error) {
	res, err := r.client.
		ProjectTaskColumn.UpdateOneID(input.ID).
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

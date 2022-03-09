package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type projectRepository struct {
	client *ent.Client
}

// NewProjectRepository generates project repository
func NewProjectRepository(client *ent.Client) ur.Project {
	return &projectRepository{client: client}
}

func (r *projectRepository) Get(ctx context.Context, where *model.ProjectWhereInput) (*model.Project, error) {
	q := r.client.Project.Query()

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

func (r *projectRepository) List(ctx context.Context) ([]*model.Project, error) {
	res, err := r.client.Project.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectWhereInput) (*model.ProjectConnection, error) {
	q := r.client.Project.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithProjectFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *projectRepository) Create(ctx context.Context, input model.CreateProjectInput) (*model.Project, error) {
	res, err := r.client.
		Project.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectRepository) Update(ctx context.Context, input model.UpdateProjectInput) (*model.Project, error) {
	res, err := r.client.
		Project.UpdateOneID(input.ID).
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

// WithProject eager-loads association with project.
func WithProject(pq *ent.ProjectQuery) {
	pq.WithProjectTeammates()
}

package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
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

	p, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, nil)
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return p, nil
}

func (r *projectRepository) List(ctx context.Context) ([]*model.Project, error) {
	ps, err := r.client.Project.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return ps, nil
}

func (r *projectRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectWhereInput, requestedFields []string) (*model.ProjectConnection, error) {
	q := r.client.Project.Query()

	if collection.Contains(requestedFields, "edges.node.icon") {
		q.WithIcon()
	}
	if collection.Contains(requestedFields, "edges.node.baseColor") {
		q.WithProjectBaseColor()
	}
	if collection.Contains(requestedFields, "edges.node.projectTeammates") {
		q.WithProjectTeammates(func(query *ent.ProjectTeammateQuery) {
			query.WithTeammate()
		})
	}

	ps, err := q.Paginate(ctx, after, first, before, last, ent.WithProjectFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return ps, nil
}

func (r *projectRepository) Create(ctx context.Context, input model.CreateProjectInput) (*model.Project, error) {
	p, err := r.client.
		Project.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return p, nil
}

func (r *projectRepository) Update(ctx context.Context, input model.UpdateProjectInput) (*model.Project, error) {
	p, err := r.client.
		Project.UpdateOneID(input.ID).
		SetInput(input).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}

		return nil, model.NewDBError(err)
	}

	return p, nil
}

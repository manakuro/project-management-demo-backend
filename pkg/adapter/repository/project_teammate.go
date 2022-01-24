package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type projectTeammateRepository struct {
	client *ent.Client
}

// NewProjectTeammateRepository generates projectTeammate repository
func NewProjectTeammateRepository(client *ent.Client) ur.ProjectTeammate {
	return &projectTeammateRepository{client: client}
}

func (r *projectTeammateRepository) Get(ctx context.Context, where *model.ProjectTeammateWhereInput) (*model.ProjectTeammate, error) {
	q := r.client.ProjectTeammate.Query()

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

func (r *projectTeammateRepository) List(ctx context.Context) ([]*model.ProjectTeammate, error) {
	res, err := r.client.
		ProjectTeammate.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectTeammateRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectTeammateWhereInput, requestedFields []string) (*model.ProjectTeammateConnection, error) {
	q := r.client.ProjectTeammate.Query()

	if collection.Contains(requestedFields, "edges.node.project") {
		q.WithProject()
	}

	if collection.Contains(requestedFields, "edges.node.teammate") {
		q.WithTeammate()
	}

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithProjectTeammateFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *projectTeammateRepository) Create(ctx context.Context, input model.CreateProjectTeammateInput) (*model.ProjectTeammate, error) {
	res, err := r.client.
		ProjectTeammate.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectTeammateRepository) Update(ctx context.Context, input model.UpdateProjectTeammateInput) (*model.ProjectTeammate, error) {
	res, err := r.client.
		ProjectTeammate.UpdateOneID(input.ID).
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

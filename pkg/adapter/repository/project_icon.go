package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type projectIconRepository struct {
	client *ent.Client
}

// NewProjectIconRepository generates projectIcon repository
func NewProjectIconRepository(client *ent.Client) ur.ProjectIcon {
	return &projectIconRepository{client: client}
}

func (r *projectIconRepository) Get(ctx context.Context, where *model.ProjectIconWhereInput) (*model.ProjectIcon, error) {
	q := r.client.ProjectIcon.Query()

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

func (r *projectIconRepository) List(ctx context.Context) ([]*model.ProjectIcon, error) {
	res, err := r.client.ProjectIcon.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectIconRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectIconWhereInput, requestedFields []string) (*model.ProjectIconConnection, error) {
	q := r.client.ProjectIcon.Query()

	if collection.Contains(requestedFields, "edges.node.icon") {
		q.WithIcon()
	}

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithProjectIconFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *projectIconRepository) Create(ctx context.Context, input model.CreateProjectIconInput) (*model.ProjectIcon, error) {
	res, err := r.client.
		ProjectIcon.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *projectIconRepository) Update(ctx context.Context, input model.UpdateProjectIconInput) (*model.ProjectIcon, error) {
	res, err := r.client.
		ProjectIcon.
		UpdateOneID(input.ID).
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

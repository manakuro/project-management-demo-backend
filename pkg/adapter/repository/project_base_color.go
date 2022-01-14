package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type projectBaseColorRepository struct {
	client *ent.Client
}

// NewProjectBaseColorRepository generates projectBaseColor repository
func NewProjectBaseColorRepository(client *ent.Client) ur.ProjectBaseColor {
	return &projectBaseColorRepository{client: client}
}

func (p *projectBaseColorRepository) Get(ctx context.Context, where *model.ProjectBaseColorWhereInput) (*model.ProjectBaseColor, error) {
	q := p.client.ProjectBaseColor.Query()

	q, err := where.Filter(q)
	if err != nil {
		return nil, model.NewInvalidParamError(nil)
	}

	result, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, nil)
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return result, nil
}

func (p *projectBaseColorRepository) List(ctx context.Context) ([]*model.ProjectBaseColor, error) {
	ps, err := p.client.
		ProjectBaseColor.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return ps, nil
}

func (p *projectBaseColorRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectBaseColorWhereInput, requestedFields []string) (*model.ProjectBaseColorConnection, error) {
	q := p.client.ProjectBaseColor.Query()

	if collection.Contains(requestedFields, "edges.node.color") {
		q.WithColor()
	}

	us, err := q.Paginate(ctx, after, first, before, last, ent.WithProjectBaseColorFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return us, nil
}

func (p *projectBaseColorRepository) Create(ctx context.Context, input model.CreateProjectBaseColorInput) (*model.ProjectBaseColor, error) {
	u, err := p.client.
		ProjectBaseColor.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return u, nil
}

func (p *projectBaseColorRepository) Update(ctx context.Context, input model.UpdateProjectBaseColorInput) (*model.ProjectBaseColor, error) {
	u, err := p.client.
		ProjectBaseColor.
		UpdateOneID(input.ID).
		SetInput(input).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input.ID)
		}

		return nil, model.NewDBError(err)
	}

	return u, nil
}

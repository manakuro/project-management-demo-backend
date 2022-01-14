package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type projectLightColorRepository struct {
	client *ent.Client
}

// NewProjectLightColorRepository generates projectLightColor repository
func NewProjectLightColorRepository(client *ent.Client) ur.ProjectLightColor {
	return &projectLightColorRepository{client: client}
}

func (p *projectLightColorRepository) Get(ctx context.Context, where *model.ProjectLightColorWhereInput) (*model.ProjectLightColor, error) {
	q := p.client.ProjectLightColor.Query()

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

func (p *projectLightColorRepository) List(ctx context.Context) ([]*model.ProjectLightColor, error) {
	ps, err := p.client.ProjectLightColor.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return ps, nil
}

func (p *projectLightColorRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectLightColorWhereInput, requestedFields []string) (*model.ProjectLightColorConnection, error) {
	q := p.client.ProjectLightColor.Query()

	if collection.Contains(requestedFields, "edges.node.color") {
		q.WithColor()
	}

	us, err := q.Paginate(ctx, after, first, before, last, ent.WithProjectLightColorFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return us, nil
}

func (p *projectLightColorRepository) Create(ctx context.Context, input model.CreateProjectLightColorInput) (*model.ProjectLightColor, error) {
	u, err := p.client.
		ProjectLightColor.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return u, nil
}

func (p *projectLightColorRepository) Update(ctx context.Context, input model.UpdateProjectLightColorInput) (*model.ProjectLightColor, error) {
	u, err := p.client.
		ProjectLightColor.
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

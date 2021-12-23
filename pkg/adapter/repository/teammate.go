package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type teammateRepository struct {
	client *ent.Client
}

// NewTeammateRepository generates teammate repository
func NewTeammateRepository(client *ent.Client) ur.Teammate {
	return &teammateRepository{client: client}
}

func (r *teammateRepository) Get(ctx context.Context, id model.ID) (*model.Teammate, error) {
	q := r.client.Teammate.Query()

	if id == "" {
		return nil, model.NewInvalidParamError(map[string]interface{}{
			"id": id,
		})
	}
	q.Where(teammate.IDEQ(id))

	u, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotSingular(err) {
			return nil, model.NewNotFoundError(err, map[string]interface{}{
				"id": id,
			})
		}
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, model.NewDBError(err)
	}

	return u, nil
}

func (r *teammateRepository) List(ctx context.Context) ([]*model.Teammate, error) {
	us, err := r.client.
		Teammate.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return us, nil
}

func (r *teammateRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TeammateWhereInput) (*model.TeammateConnection, error) {
	us, err := r.client.
		Teammate.
		Query().
		Paginate(ctx, after, first, before, last, ent.WithTeammateFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return us, nil
}

func (r *teammateRepository) Create(ctx context.Context, input model.CreateTeammateInput) (*model.Teammate, error) {
	u, err := r.client.
		Teammate.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return u, nil
}

func (r *teammateRepository) Update(ctx context.Context, input model.UpdateTeammateInput) (*model.Teammate, error) {
	u, err := r.client.
		Teammate.UpdateOneID(input.ID).
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

package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/color"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type colorRepository struct {
	client *ent.Client
}

// NewColorRepository generates color repository
func NewColorRepository(client *ent.Client) ur.Color {
	return &colorRepository{client: client}
}

func (r *colorRepository) Get(ctx context.Context, id model.ID) (*model.Color, error) {
	q := r.client.Color.Query()

	if id == "" {
		return nil, model.NewInvalidParamError(map[string]interface{}{
			"id": id,
		})
	}
	q.Where(color.IDEQ(id))

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

func (r *colorRepository) List(ctx context.Context) ([]*model.Color, error) {
	us, err := r.client.
		Color.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return us, nil
}

func (r *colorRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ColorWhereInput) (*model.ColorConnection, error) {
	us, err := r.client.
		Color.
		Query().
		Paginate(ctx, after, first, before, last, ent.WithColorFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return us, nil
}

func (r *colorRepository) Create(ctx context.Context, input model.CreateColorInput) (*model.Color, error) {
	u, err := r.client.
		Color.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return u, nil
}

func (r *colorRepository) Update(ctx context.Context, input model.UpdateColorInput) (*model.Color, error) {
	u, err := r.client.
		Color.UpdateOneID(input.ID).
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

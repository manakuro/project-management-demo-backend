package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/icon"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type iconRepository struct {
	client *ent.Client
}

// NewIconRepository generates icon repository
func NewIconRepository(client *ent.Client) ur.Icon {
	return &iconRepository{client: client}
}

func (r *iconRepository) Get(ctx context.Context, id model.ID) (*model.Icon, error) {
	q := r.client.Icon.Query()

	if id == "" {
		return nil, model.NewInvalidParamError(map[string]interface{}{
			"id": id,
		})
	}
	q.Where(icon.IDEQ(id))

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

func (r *iconRepository) List(ctx context.Context) ([]*model.Icon, error) {
	us, err := r.client.
		Icon.
		Query().
		All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return us, nil
}

func (r *iconRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.IconWhereInput) (*model.IconConnection, error) {
	us, err := r.client.
		Icon.
		Query().
		Paginate(ctx, after, first, before, last, ent.WithIconFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return us, nil
}

func (r *iconRepository) Create(ctx context.Context, input model.CreateIconInput) (*model.Icon, error) {
	u, err := r.client.
		Icon.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return u, nil
}

func (r *iconRepository) Update(ctx context.Context, input model.UpdateIconInput) (*model.Icon, error) {
	u, err := r.client.
		Icon.UpdateOneID(input.ID).
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

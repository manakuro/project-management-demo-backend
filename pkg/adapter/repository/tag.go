package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type tagRepository struct {
	client *ent.Client
}

// NewTagRepository generates tag repository
func NewTagRepository(client *ent.Client) ur.Tag {
	return &tagRepository{client: client}
}

func (r *tagRepository) Get(ctx context.Context, where *model.TagWhereInput) (*model.Tag, error) {
	q := r.client.Tag.Query()

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

func (r *tagRepository) List(ctx context.Context) ([]*model.Tag, error) {
	res, err := r.client.Tag.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *tagRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TagWhereInput) (*model.TagConnection, error) {
	q := r.client.Tag.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithTagFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *tagRepository) Create(ctx context.Context, input model.CreateTagInput) (*model.Tag, error) {
	res, err := r.client.
		Tag.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *tagRepository) Update(ctx context.Context, input model.UpdateTagInput) (*model.Tag, error) {
	res, err := r.client.
		Tag.UpdateOneID(input.ID).
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

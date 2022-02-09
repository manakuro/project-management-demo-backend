package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/favoriteproject"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
)

type favoriteProjectRepository struct {
	client *ent.Client
}

// NewFavoriteProjectRepository generates favoriteProject repository
func NewFavoriteProjectRepository(client *ent.Client) ur.FavoriteProject {
	return &favoriteProjectRepository{client: client}
}

func (r *favoriteProjectRepository) Get(ctx context.Context, where *model.FavoriteProjectWhereInput) (*model.FavoriteProject, error) {
	q := r.client.FavoriteProject.Query()

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

func (r *favoriteProjectRepository) List(ctx context.Context) ([]*model.FavoriteProject, error) {
	res, err := r.client.
		FavoriteProject.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *favoriteProjectRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.FavoriteProjectWhereInput) (*model.FavoriteProjectConnection, error) {
	q := r.client.FavoriteProject.Query()

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithFavoriteProjectFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *favoriteProjectRepository) Create(ctx context.Context, input model.CreateFavoriteProjectInput) (*model.FavoriteProject, error) {
	res, err := r.client.
		FavoriteProject.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *favoriteProjectRepository) Update(ctx context.Context, input model.UpdateFavoriteProjectInput) (*model.FavoriteProject, error) {
	res, err := r.client.
		FavoriteProject.UpdateOneID(input.ID).
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

func (r *favoriteProjectRepository) Delete(ctx context.Context, input model.DeleteFavoriteProjectInput) (*model.FavoriteProject, error) {
	deleted, err := r.client.
		FavoriteProject.
		Query().
		Where(favoriteproject.ProjectID(input.ProjectID)).
		Where(favoriteproject.TeammateID(input.TeammateID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	_, err = r.client.
		FavoriteProject.
		Delete().
		Where(favoriteproject.ProjectID(input.ProjectID)).
		Where(favoriteproject.TeammateID(input.TeammateID)).
		Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}

func (r *favoriteProjectRepository) FavoriteProjectIDs(ctx context.Context, teammateID model.ID) ([]model.ID, error) {
	res, err := r.client.FavoriteProject.Query().Where(favoriteproject.TeammateID(teammateID)).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}

		return nil, model.NewDBError(err)
	}

	ids := make([]model.ID, len(res))
	for i, fp := range res {
		ids[i] = fp.ProjectID
	}

	return ids, nil
}

package repository

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/favoriteworkspace"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"
	"project-management-demo-backend/pkg/util/collection"
)

type favoriteWorkspaceRepository struct {
	client *ent.Client
}

// NewFavoriteWorkspaceRepository generates favoriteWorkspace repository
func NewFavoriteWorkspaceRepository(client *ent.Client) ur.FavoriteWorkspace {
	return &favoriteWorkspaceRepository{client: client}
}

func (r *favoriteWorkspaceRepository) Get(ctx context.Context, where *model.FavoriteWorkspaceWhereInput) (*model.FavoriteWorkspace, error) {
	q := r.client.FavoriteWorkspace.Query()

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

func (r *favoriteWorkspaceRepository) List(ctx context.Context) ([]*model.FavoriteWorkspace, error) {
	res, err := r.client.
		FavoriteWorkspace.Query().All(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *favoriteWorkspaceRepository) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.FavoriteWorkspaceWhereInput, requestedFields []string) (*model.FavoriteWorkspaceConnection, error) {
	q := r.client.FavoriteWorkspace.Query()

	if collection.Contains(requestedFields, "edges.node.project") {
		q.WithWorkspace()
	}

	if collection.Contains(requestedFields, "edges.node.teammate") {
		q.WithTeammate()
	}

	res, err := q.Paginate(ctx, after, first, before, last, ent.WithFavoriteWorkspaceFilter(where.Filter))
	if err != nil {
		return nil, model.NewDBError(err)
	}
	return res, nil
}

func (r *favoriteWorkspaceRepository) Create(ctx context.Context, input model.CreateFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error) {
	res, err := r.client.
		FavoriteWorkspace.
		Create().
		SetInput(input).
		Save(ctx)

	if err != nil {
		return nil, model.NewDBError(err)
	}

	return res, nil
}

func (r *favoriteWorkspaceRepository) Update(ctx context.Context, input model.UpdateFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error) {
	res, err := r.client.
		FavoriteWorkspace.UpdateOneID(input.ID).
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

func (r *favoriteWorkspaceRepository) Delete(ctx context.Context, input model.DeleteFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error) {
	deleted, err := r.client.
		FavoriteWorkspace.
		Query().
		Where(favoriteworkspace.WorkspaceID(input.WorkspaceID)).
		Where(favoriteworkspace.TeammateID(input.TeammateID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, model.NewNotFoundError(err, input)
		}
		return nil, model.NewDBError(err)
	}

	_, err = r.client.
		FavoriteWorkspace.
		Delete().
		Where(favoriteworkspace.WorkspaceID(input.WorkspaceID)).
		Where(favoriteworkspace.TeammateID(input.TeammateID)).
		Exec(ctx)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	return deleted, nil
}

func (r *favoriteWorkspaceRepository) FavoriteWorkspaceIDs(ctx context.Context, teammateID model.ID, workspaceID *model.ID) ([]model.ID, error) {
	q := r.client.
		FavoriteWorkspace.
		Query()

	if workspaceID != nil {
		q.Where(favoriteworkspace.WorkspaceID(*workspaceID))
	}

	res, err := q.Where(favoriteworkspace.TeammateID(teammateID)).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}

		return nil, model.NewDBError(err)
	}

	ids := make([]model.ID, len(res))
	for i, fp := range res {
		ids[i] = fp.WorkspaceID
	}

	return ids, nil
}

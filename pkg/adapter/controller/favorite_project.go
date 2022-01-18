package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// FavoriteProject is an interface of controller
type FavoriteProject interface {
	Get(ctx context.Context, where *model.FavoriteProjectWhereInput) (*model.FavoriteProject, error)
	List(ctx context.Context) ([]*model.FavoriteProject, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.FavoriteProjectWhereInput, requestedFields []string) (*model.FavoriteProjectConnection, error)
	Create(ctx context.Context, input model.CreateFavoriteProjectInput) (*model.FavoriteProject, error)
	Update(ctx context.Context, input model.UpdateFavoriteProjectInput) (*model.FavoriteProject, error)
	FavoriteProjectIDs(ctx context.Context, teammateID model.ID) ([]model.ID, error)
}

type favoriteProject struct {
	favoriteProjectUsecase usecase.FavoriteProject
}

// NewFavoriteProjectController generates favoriteProject controller
func NewFavoriteProjectController(pt usecase.FavoriteProject) FavoriteProject {
	return &favoriteProject{
		favoriteProjectUsecase: pt,
	}
}

func (f *favoriteProject) Get(ctx context.Context, where *model.FavoriteProjectWhereInput) (*model.FavoriteProject, error) {
	return f.favoriteProjectUsecase.Get(ctx, where)
}

func (f *favoriteProject) List(ctx context.Context) ([]*model.FavoriteProject, error) {
	return f.favoriteProjectUsecase.List(ctx)
}

func (f *favoriteProject) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.FavoriteProjectWhereInput, requestedFields []string) (*model.FavoriteProjectConnection, error) {
	return f.favoriteProjectUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (f *favoriteProject) Create(ctx context.Context, input model.CreateFavoriteProjectInput) (*model.FavoriteProject, error) {
	return f.favoriteProjectUsecase.Create(ctx, input)
}

func (f *favoriteProject) Update(ctx context.Context, input model.UpdateFavoriteProjectInput) (*model.FavoriteProject, error) {
	return f.favoriteProjectUsecase.Update(ctx, input)
}

func (f *favoriteProject) FavoriteProjectIDs(ctx context.Context, teammateID model.ID) ([]model.ID, error) {
	return f.favoriteProjectUsecase.FavoriteProjectIDs(ctx, teammateID)
}

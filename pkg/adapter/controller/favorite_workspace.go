package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// FavoriteWorkspace is an interface of controller
type FavoriteWorkspace interface {
	Get(ctx context.Context, where *model.FavoriteWorkspaceWhereInput) (*model.FavoriteWorkspace, error)
	List(ctx context.Context) ([]*model.FavoriteWorkspace, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.FavoriteWorkspaceWhereInput, requestedFields []string) (*model.FavoriteWorkspaceConnection, error)
	Create(ctx context.Context, input model.CreateFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error)
	Update(ctx context.Context, input model.UpdateFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error)
	Delete(ctx context.Context, input model.DeleteFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error)
	FavoriteWorkspaceIDs(ctx context.Context, teammateID model.ID, workspaceID *model.ID) ([]model.ID, error)
}

type favoriteWorkspace struct {
	favoriteWorkspaceUsecase usecase.FavoriteWorkspace
}

// NewFavoriteWorkspaceController generates favoriteWorkspace controller
func NewFavoriteWorkspaceController(pt usecase.FavoriteWorkspace) FavoriteWorkspace {
	return &favoriteWorkspace{
		favoriteWorkspaceUsecase: pt,
	}
}

func (f *favoriteWorkspace) Get(ctx context.Context, where *model.FavoriteWorkspaceWhereInput) (*model.FavoriteWorkspace, error) {
	return f.favoriteWorkspaceUsecase.Get(ctx, where)
}

func (f *favoriteWorkspace) List(ctx context.Context) ([]*model.FavoriteWorkspace, error) {
	return f.favoriteWorkspaceUsecase.List(ctx)
}

func (f *favoriteWorkspace) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.FavoriteWorkspaceWhereInput, requestedFields []string) (*model.FavoriteWorkspaceConnection, error) {
	return f.favoriteWorkspaceUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (f *favoriteWorkspace) Create(ctx context.Context, input model.CreateFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error) {
	return f.favoriteWorkspaceUsecase.Create(ctx, input)
}

func (f *favoriteWorkspace) Update(ctx context.Context, input model.UpdateFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error) {
	return f.favoriteWorkspaceUsecase.Update(ctx, input)
}

func (f *favoriteWorkspace) Delete(ctx context.Context, input model.DeleteFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error) {
	return f.favoriteWorkspaceUsecase.Delete(ctx, input)
}

func (f *favoriteWorkspace) FavoriteWorkspaceIDs(ctx context.Context, teammateID model.ID, workspaceID *model.ID) ([]model.ID, error) {
	return f.favoriteWorkspaceUsecase.FavoriteWorkspaceIDs(ctx, teammateID, workspaceID)
}

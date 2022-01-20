package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type favoriteWorkspace struct {
	favoriteWorkspaceRepository repository.FavoriteWorkspace
}

// FavoriteWorkspace is an interface of test user
type FavoriteWorkspace interface {
	Get(ctx context.Context, where *model.FavoriteWorkspaceWhereInput) (*model.FavoriteWorkspace, error)
	List(ctx context.Context) ([]*model.FavoriteWorkspace, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.FavoriteWorkspaceWhereInput, requestedFields []string) (*model.FavoriteWorkspaceConnection, error)
	Create(ctx context.Context, input model.CreateFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error)
	Update(ctx context.Context, input model.UpdateFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error)
	Delete(ctx context.Context, input model.DeleteFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error)
	FavoriteWorkspaceIDs(ctx context.Context, teammateID model.ID, workspaceID *model.ID) ([]model.ID, error)
}

// NewFavoriteWorkspaceUsecase generates test user repository
func NewFavoriteWorkspaceUsecase(r repository.FavoriteWorkspace) FavoriteWorkspace {
	return &favoriteWorkspace{favoriteWorkspaceRepository: r}
}

func (f *favoriteWorkspace) Get(ctx context.Context, where *model.FavoriteWorkspaceWhereInput) (*model.FavoriteWorkspace, error) {
	return f.favoriteWorkspaceRepository.Get(ctx, where)
}

func (f *favoriteWorkspace) List(ctx context.Context) ([]*model.FavoriteWorkspace, error) {
	return f.favoriteWorkspaceRepository.List(ctx)
}

func (f *favoriteWorkspace) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.FavoriteWorkspaceWhereInput, requestedFields []string) (*model.FavoriteWorkspaceConnection, error) {
	return f.favoriteWorkspaceRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (f *favoriteWorkspace) Create(ctx context.Context, input model.CreateFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error) {
	return f.favoriteWorkspaceRepository.Create(ctx, input)
}

func (f *favoriteWorkspace) Update(ctx context.Context, input model.UpdateFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error) {
	return f.favoriteWorkspaceRepository.Update(ctx, input)
}

func (f *favoriteWorkspace) Delete(ctx context.Context, input model.DeleteFavoriteWorkspaceInput) (*model.FavoriteWorkspace, error) {
	return f.favoriteWorkspaceRepository.Delete(ctx, input)
}

func (f *favoriteWorkspace) FavoriteWorkspaceIDs(ctx context.Context, teammateID model.ID, workspaceID *model.ID) ([]model.ID, error) {
	return f.favoriteWorkspaceRepository.FavoriteWorkspaceIDs(ctx, teammateID, workspaceID)
}

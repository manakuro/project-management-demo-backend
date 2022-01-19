package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type favoriteProject struct {
	favoriteProjectRepository repository.FavoriteProject
}

// FavoriteProject is an interface of test user
type FavoriteProject interface {
	Get(ctx context.Context, where *model.FavoriteProjectWhereInput) (*model.FavoriteProject, error)
	List(ctx context.Context) ([]*model.FavoriteProject, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.FavoriteProjectWhereInput, requestedFields []string) (*model.FavoriteProjectConnection, error)
	Create(ctx context.Context, input model.CreateFavoriteProjectInput) (*model.FavoriteProject, error)
	Update(ctx context.Context, input model.UpdateFavoriteProjectInput) (*model.FavoriteProject, error)
	Delete(ctx context.Context, input model.DeleteFavoriteProjectInput) (int, error)
	FavoriteProjectIDs(ctx context.Context, teammateID model.ID) ([]model.ID, error)
}

// NewFavoriteProjectUsecase generates test user repository
func NewFavoriteProjectUsecase(r repository.FavoriteProject) FavoriteProject {
	return &favoriteProject{favoriteProjectRepository: r}
}

func (f *favoriteProject) Get(ctx context.Context, where *model.FavoriteProjectWhereInput) (*model.FavoriteProject, error) {
	return f.favoriteProjectRepository.Get(ctx, where)
}

func (f *favoriteProject) List(ctx context.Context) ([]*model.FavoriteProject, error) {
	return f.favoriteProjectRepository.List(ctx)
}

func (f *favoriteProject) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.FavoriteProjectWhereInput, requestedFields []string) (*model.FavoriteProjectConnection, error) {
	return f.favoriteProjectRepository.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (f *favoriteProject) Create(ctx context.Context, input model.CreateFavoriteProjectInput) (*model.FavoriteProject, error) {
	return f.favoriteProjectRepository.Create(ctx, input)
}

func (f *favoriteProject) Update(ctx context.Context, input model.UpdateFavoriteProjectInput) (*model.FavoriteProject, error) {
	return f.favoriteProjectRepository.Update(ctx, input)
}

func (f *favoriteProject) Delete(ctx context.Context, input model.DeleteFavoriteProjectInput) (int, error) {
	return f.favoriteProjectRepository.Delete(ctx, input)
}

func (f *favoriteProject) FavoriteProjectIDs(ctx context.Context, teammateID model.ID) ([]model.ID, error) {
	return f.favoriteProjectRepository.FavoriteProjectIDs(ctx, teammateID)
}

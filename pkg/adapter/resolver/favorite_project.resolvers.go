package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/handler"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/util/datetime"
	"project-management-demo-backend/pkg/util/graphqlutil"
)

func (r *favoriteProjectResolver) CreatedAt(ctx context.Context, obj *ent.FavoriteProject) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *favoriteProjectResolver) UpdatedAt(ctx context.Context, obj *ent.FavoriteProject) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

func (r *mutationResolver) CreateFavoriteProject(ctx context.Context, input ent.CreateFavoriteProjectInput) (*ent.FavoriteProject, error) {
	f, err := r.controller.FavoriteProject.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return f, nil
}

func (r *mutationResolver) DeleteFavoriteProject(ctx context.Context, input model.DeleteFavoriteProjectInput) (int, error) {
	result, err := r.controller.FavoriteProject.Delete(ctx, input)
	if err != nil {
		return 0, handler.HandleGraphQLError(ctx, err)
	}

	return result, nil
}

func (r *queryResolver) FavoriteProject(ctx context.Context, where *ent.FavoriteProjectWhereInput) (*ent.FavoriteProject, error) {
	f, err := r.controller.FavoriteProject.Get(ctx, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return f, nil
}

func (r *queryResolver) FavoriteProjects(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.FavoriteProjectWhereInput) (*ent.FavoriteProjectConnection, error) {
	requestedFields := graphqlutil.GetRequestedFields(ctx)

	fs, err := r.controller.FavoriteProject.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return fs, nil
}

func (r *queryResolver) FavoriteProjectIds(ctx context.Context, teammateID ulid.ID) ([]ulid.ID, error) {
	ids, err := r.controller.FavoriteProject.FavoriteProjectIDs(ctx, teammateID)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return ids, nil
}

// FavoriteProject returns generated.FavoriteProjectResolver implementation.
func (r *Resolver) FavoriteProject() generated.FavoriteProjectResolver {
	return &favoriteProjectResolver{r}
}

type favoriteProjectResolver struct{ *Resolver }

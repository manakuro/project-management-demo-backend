package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/handler"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/util/datetime"
	"project-management-demo-backend/pkg/util/graphqlutil"
	"project-management-demo-backend/pkg/util/subscription"
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

	go func() {
		ids, _ := r.controller.FavoriteProject.FavoriteProjectIDs(context.Background(), input.TeammateID)
		for _, created := range r.subscriptions.FavoriteProjectIDsUpdated {
			if created.TeammateID == input.TeammateID {
				created.Ch <- ids
			}
		}
	}()

	return f, nil
}

func (r *mutationResolver) DeleteFavoriteProject(ctx context.Context, input model.DeleteFavoriteProjectInput) (*ent.FavoriteProject, error) {
	result, err := r.controller.FavoriteProject.Delete(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	go func() {
		ids, _ := r.controller.FavoriteProject.FavoriteProjectIDs(context.Background(), input.TeammateID)
		for _, created := range r.subscriptions.FavoriteProjectIDsUpdated {
			if created.TeammateID == input.TeammateID {
				created.Ch <- ids
			}
		}
	}()

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

func (r *subscriptionResolver) FavoriteProjectCreated(ctx context.Context, teammateID ulid.ID) (<-chan *ent.FavoriteProject, error) {
	key := subscription.NewKey()
	ch := make(chan *ent.FavoriteProject, 1)

	r.mutex.Lock()
	r.subscriptions.FavoriteProjectCreated[key] = subscription.FavoriteProjectCreated{
		TeammateID: teammateID,
		Ch:         ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.FavoriteProjectCreated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

func (r *subscriptionResolver) FavoriteProjectIdsUpdated(ctx context.Context, teammateID ulid.ID) (<-chan []ulid.ID, error) {
	key := subscription.NewKey()
	ch := make(chan []ulid.ID, 1)

	r.mutex.Lock()
	r.subscriptions.FavoriteProjectIDsUpdated[key] = subscription.FavoriteProjectIDsUpdated{
		TeammateID: teammateID,
		Ch:         ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.FavoriteProjectIDsUpdated, key)
		fmt.Println("FavoriteProjectIdsUpdated: finished!")
		r.mutex.Unlock()
	}()

	return ch, nil
}

// FavoriteProject returns generated.FavoriteProjectResolver implementation.
func (r *Resolver) FavoriteProject() generated.FavoriteProjectResolver {
	return &favoriteProjectResolver{r}
}

type favoriteProjectResolver struct{ *Resolver }

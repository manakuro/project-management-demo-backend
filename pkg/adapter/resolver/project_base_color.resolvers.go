package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/handler"
	"project-management-demo-backend/pkg/util/datetime"
	"project-management-demo-backend/pkg/util/subscription"
)

func (r *mutationResolver) CreateProjectBaseColor(ctx context.Context, input ent.CreateProjectBaseColorInput) (*ent.ProjectBaseColor, error) {
	p, err := r.controller.ProjectBaseColor.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return p, nil
}

func (r *mutationResolver) UpdateProjectBaseColor(ctx context.Context, input ent.UpdateProjectBaseColorInput) (*ent.ProjectBaseColor, error) {
	p, err := r.controller.ProjectBaseColor.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	for _, pu := range r.subscriptions.ProjectBaseColorUpdated {
		if pu.ID == p.ID {
			pu.Ch <- p
		}
	}

	return p, nil
}

func (r *projectBaseColorResolver) CreatedAt(ctx context.Context, obj *ent.ProjectBaseColor) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *projectBaseColorResolver) UpdatedAt(ctx context.Context, obj *ent.ProjectBaseColor) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

func (r *queryResolver) ProjectBaseColor(ctx context.Context, where *ent.ProjectBaseColorWhereInput) (*ent.ProjectBaseColor, error) {
	p, err := r.controller.ProjectBaseColor.Get(ctx, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return p, nil
}

func (r *queryResolver) ProjectBaseColors(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ProjectBaseColorWhereInput) (*ent.ProjectBaseColorConnection, error) {
	ps, err := r.controller.ProjectBaseColor.ListWithPagination(ctx, after, first, before, last, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return ps, nil
}

func (r *subscriptionResolver) ProjectBaseColorUpdated(ctx context.Context, id ulid.ID) (<-chan *ent.ProjectBaseColor, error) {
	key := subscription.NewKey()
	ch := make(chan *ent.ProjectBaseColor, 1)

	r.mutex.Lock()
	r.subscriptions.ProjectBaseColorUpdated[key] = subscription.ProjectBaseColorUpdated{
		ID: id,
		Ch: ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.ProjectBaseColorUpdated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

// ProjectBaseColor returns generated.ProjectBaseColorResolver implementation.
func (r *Resolver) ProjectBaseColor() generated.ProjectBaseColorResolver {
	return &projectBaseColorResolver{r}
}

type projectBaseColorResolver struct{ *Resolver }

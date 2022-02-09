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

func (r *mutationResolver) CreateTeammateTaskSection(ctx context.Context, input ent.CreateTeammateTaskSectionInput) (*ent.TeammateTaskSection, error) {
	t, err := r.controller.TeammateTaskSection.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return t, nil
}

func (r *mutationResolver) UpdateTeammateTaskSection(ctx context.Context, input ent.UpdateTeammateTaskSectionInput) (*ent.TeammateTaskSection, error) {
	t, err := r.controller.TeammateTaskSection.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	for _, u := range r.subscriptions.TeammateTaskSectionUpdated {
		if u.ID == t.ID {
			u.Ch <- t
		}
	}

	return t, nil
}

func (r *queryResolver) TeammateTaskSection(ctx context.Context, where *ent.TeammateTaskSectionWhereInput) (*ent.TeammateTaskSection, error) {
	t, err := r.controller.TeammateTaskSection.Get(ctx, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}
	return t, nil
}

func (r *queryResolver) TeammateTaskSections(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.TeammateTaskSectionWhereInput) (*ent.TeammateTaskSectionConnection, error) {

	ts, err := r.controller.TeammateTaskSection.ListWithPagination(ctx, after, first, before, last, where)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return ts, nil
}

func (r *subscriptionResolver) TeammateTaskSectionUpdated(ctx context.Context, id ulid.ID) (<-chan *ent.TeammateTaskSection, error) {
	key := subscription.NewKey()
	ch := make(chan *ent.TeammateTaskSection, 1)

	r.mutex.Lock()
	r.subscriptions.TeammateTaskSectionUpdated[key] = subscription.TeammateTaskSectionUpdated{
		ID: id,
		Ch: ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.TeammateTaskSectionUpdated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

func (r *teammateTaskSectionResolver) CreatedAt(ctx context.Context, obj *ent.TeammateTaskSection) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *teammateTaskSectionResolver) UpdatedAt(ctx context.Context, obj *ent.TeammateTaskSection) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// TeammateTaskSection returns generated.TeammateTaskSectionResolver implementation.
func (r *Resolver) TeammateTaskSection() generated.TeammateTaskSectionResolver {
	return &teammateTaskSectionResolver{r}
}

type teammateTaskSectionResolver struct{ *Resolver }

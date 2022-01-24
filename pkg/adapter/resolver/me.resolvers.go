package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/handler"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/util/datetime"
	"project-management-demo-backend/pkg/util/subscription"
)

func (r *meResolver) CreatedAt(ctx context.Context, obj *model.Me) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *meResolver) UpdatedAt(ctx context.Context, obj *model.Me) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

func (r *mutationResolver) UpdateMe(ctx context.Context, input model.UpdateMeInput) (*model.Me, error) {
	me, err := r.controller.Me.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	for _, mu := range r.subscriptions.MeUpdated {
		if mu.ID == me.ID {
			mu.Ch <- me
		}
	}

	for _, tu := range r.subscriptions.TeammateUpdated {
		if tu.ID == me.ID {
			tu.Ch <- &model.Teammate{
				ID:        me.ID,
				Email:     me.Email,
				Name:      me.Name,
				CreatedAt: me.CreatedAt,
				UpdatedAt: me.UpdatedAt,
			}
		}
	}

	return me, err
}

func (r *queryResolver) Me(ctx context.Context) (*model.Me, error) {
	me, err := r.client.Teammate.Query().Where(teammate.EmailEQ("manato.kuroda@example.com")).Only(ctx)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	m, err := r.controller.Me.Get(ctx, me.ID)
	if err != nil {
		return nil, handler.HandleGraphQLError(ctx, err)
	}

	return m, nil
}

func (r *subscriptionResolver) MeUpdated(ctx context.Context, id ulid.ID) (<-chan *model.Me, error) {
	key := subscription.NewKey()
	ch := make(chan *model.Me)

	r.mutex.Lock()
	r.subscriptions.MeUpdated[key] = subscription.MeUpdated{
		ID: id,
		Ch: ch,
	}
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscriptions.MeUpdated, key)
		r.mutex.Unlock()
	}()

	return ch, nil
}

// Me returns generated.MeResolver implementation.
func (r *Resolver) Me() generated.MeResolver { return &meResolver{r} }

type meResolver struct{ *Resolver }

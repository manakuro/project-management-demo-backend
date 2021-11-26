package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/pulid"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/handler"
	"project-management-demo-backend/pkg/util/datetime"
)

func (r *mutationResolver) CreateTestUser(ctx context.Context, input ent.CreateTestUserInput) (*ent.TestUser, error) {
	u, err := r.controller.TestUser.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *mutationResolver) UpdateTestUser(ctx context.Context, input ent.UpdateTestUserInput) (*ent.TestUser, error) {
	u, err := r.controller.TestUser.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *queryResolver) TestUser(ctx context.Context, id *pulid.ID, age *int) (*ent.TestUser, error) {
	u, err := r.controller.TestUser.Get(ctx, id, age)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *queryResolver) TestUsers(ctx context.Context) ([]*ent.TestUser, error) {
	us, err := r.controller.TestUser.List(ctx)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return us, nil
}

func (r *testUserResolver) CreatedAt(ctx context.Context, obj *ent.TestUser) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *testUserResolver) UpdatedAt(ctx context.Context, obj *ent.TestUser) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

// TestUser returns generated.TestUserResolver implementation.
func (r *Resolver) TestUser() generated.TestUserResolver { return &testUserResolver{r} }

type testUserResolver struct{ *Resolver }

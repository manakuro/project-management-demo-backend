package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/handler"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *mutationResolver) CreateTestUser(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error) {
	u, err := r.controller.TestUser.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *mutationResolver) UpdateTestUser(ctx context.Context, input model.UpdateTestUserInput) (*model.TestUser, error) {
	u, err := r.controller.TestUser.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *queryResolver) TestUser(ctx context.Context, id *string, age *int) (*model.TestUser, error) {
	u, err := r.controller.TestUser.Get(ctx, id, age)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *testUserResolver) CreatedAt(ctx context.Context, obj *model.TestUser) (string, error) {
	return obj.FormattedCreatedAt(), nil
}

func (r *testUserResolver) UpdatedAt(ctx context.Context, obj *model.TestUser) (string, error) {
	return obj.FormattedUpdatedAt(), nil
}

// TestUser returns generated.TestUserResolver implementation.
func (r *Resolver) TestUser() generated.TestUserResolver { return &testUserResolver{r} }

type testUserResolver struct{ *Resolver }

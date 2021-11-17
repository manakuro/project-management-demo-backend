package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/entity/model"
	"strconv"
)

func (r *mutationResolver) CreateTestUser(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error) {
	return r.controller.TestUser.Create(ctx, input)
}

func (r *mutationResolver) UpdateTestUser(ctx context.Context, id string, input model.UpdateTestUserInput) (*model.TestUser, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if !errors.Is(err, nil) {
		return nil, err
	}
	u, err := r.client.TestUser.UpdateOneID(int(i)).SetInput(input.UpdateTestUserInput).Save(ctx)

	return &model.TestUser{TestUser: u}, err
}

func (r *queryResolver) TestUser(ctx context.Context) (*model.TestUser, error) {
	return r.controller.TestUser.Get(ctx)
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

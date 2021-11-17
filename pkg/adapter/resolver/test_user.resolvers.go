package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/entity/model"
	"time"
)

func (r *queryResolver) TestUser(ctx context.Context) (*model.TestUser, error) {
	return r.controller.TestUser.Get(ctx)
}

func (r *testUserResolver) CreatedAt(ctx context.Context, obj *model.TestUser) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

// TestUser returns generated.TestUserResolver implementation.
func (r *Resolver) TestUser() generated.TestUserResolver { return &testUserResolver{r} }

type testUserResolver struct{ *Resolver }

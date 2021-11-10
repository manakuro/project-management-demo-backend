package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/testuser"
	"project-management-demo-backend/graph/generated"
	"time"
)

func (r *queryResolver) User(ctx context.Context) (*ent.TestUser, error) {
	return r.client.TestUser.Query().Where(testuser.IDEQ(1)).Only(ctx)
}

func (r *testUserResolver) CreatedAt(_ context.Context, obj *ent.TestUser) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

// TestUser returns generated.TestUserResolver implementation.
func (r *Resolver) TestUser() generated.TestUserResolver { return &testUserResolver{r} }

type testUserResolver struct{ *Resolver }

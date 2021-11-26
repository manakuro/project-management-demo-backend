package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/pulid"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/handler"
	"project-management-demo-backend/pkg/entity/model"
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *testUserResolver) TestTodos(ctx context.Context, obj *ent.TestUser) ([]*model.TestTodo, error) {
	ts, err := obj.TestTodos(ctx)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}

	return ts, nil
}

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

func (r *queryResolver) TestUser(ctx context.Context, id *model.ID, age *int) (*model.TestUser, error) {
	u, err := r.controller.TestUser.Get(ctx, id, age)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return u, nil
}

func (r *queryResolver) TestUsers(ctx context.Context) ([]*model.TestUser, error) {
	us, err := r.controller.TestUser.List(ctx)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return us, nil
}

func (r *testUserResolver) ID(ctx context.Context, obj *model.TestUser) (model.ID, error) {
	return model.ID(obj.ID), nil
}

func (r *testUserResolver) TestTodos(ctx context.Context, obj *model.TestUser) ([]*model.TestTodo, error) {
	ts, err := obj.TestTodos(ctx)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	res := make([]*model.TestTodo, len(ts))
	for i, t := range ts {
		res[i] = &model.TestTodo{TestTodo: t}
	}

	return res, nil
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

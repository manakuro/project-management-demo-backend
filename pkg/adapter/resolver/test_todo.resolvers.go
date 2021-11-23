package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/handler"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *mutationResolver) CreateTestTodo(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error) {
	t, err := r.controller.TestTodo.Create(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return t, nil
}

func (r *mutationResolver) UpdateTestTodo(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error) {
	t, err := r.controller.TestTodo.Update(ctx, input)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return t, nil
}

func (r *queryResolver) TestTodo(ctx context.Context, id *string) (*model.TestTodo, error) {
	t, err := r.controller.TestTodo.Get(ctx, id)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return t, nil
}

func (r *queryResolver) TestTodos(ctx context.Context) ([]*model.TestTodo, error) {
	ts, err := r.controller.TestTodo.List(ctx)
	if err != nil {
		return nil, handler.HandleError(ctx, err)
	}
	return ts, nil
}

func (r *testTodoResolver) CreatedAt(ctx context.Context, obj *model.TestTodo) (string, error) {
	return obj.FormattedCreatedAt(), nil
}

func (r *testTodoResolver) UpdatedAt(ctx context.Context, obj *model.TestTodo) (string, error) {
	return obj.FormattedUpdatedAt(), nil
}

// TestTodo returns generated.TestTodoResolver implementation.
func (r *Resolver) TestTodo() generated.TestTodoResolver { return &testTodoResolver{r} }

type testTodoResolver struct{ *Resolver }

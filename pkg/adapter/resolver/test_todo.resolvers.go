package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/entity/model"
)

func (r *mutationResolver) CreateTestTodo(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTestTodo(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TestTodo(ctx context.Context, id *string) (*model.TestTodo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *testTodoResolver) CreatedAt(ctx context.Context, obj *model.TestTodo) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *testTodoResolver) UpdatedAt(ctx context.Context, obj *model.TestTodo) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// TestTodo returns generated.TestTodoResolver implementation.
func (r *Resolver) TestTodo() generated.TestTodoResolver { return &testTodoResolver{r} }

type testTodoResolver struct{ *Resolver }

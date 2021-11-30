package repository

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
)

// TestTodo is interface of repository
type TestTodo interface {
	Get(ctx context.Context, id *model.ID) (*model.TestTodo, error)
	List(ctx context.Context) ([]*model.TestTodo, error)
	Create(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error)
	Update(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error)
}

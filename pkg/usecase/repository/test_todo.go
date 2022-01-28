package repository

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
)

// TestTodo is interface of repository
type TestTodo interface {
	Get(ctx context.Context, where *model.TestTodoWhereInput) (*model.TestTodo, error)
	List(ctx context.Context) ([]*model.TestTodo, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TestTodoWhereInput, requestFields []string) (*model.TestTodoConnection, error)
	Create(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error)
	Update(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error)
}

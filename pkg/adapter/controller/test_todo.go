package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TestTodo is an interface of controller.
type TestTodo interface {
	Get(ctx context.Context, where *model.TestTodoWhereInput) (*model.TestTodo, error)
	List(ctx context.Context) ([]*model.TestTodo, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TestTodoWhereInput, requestedFields []string) (*model.TestTodoConnection, error)
	Create(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error)
	Update(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error)
}

type testTodoController struct {
	testTodoUsecase usecase.TestTodo
}

// NewTestTodoController generates controller.
func NewTestTodoController(pt usecase.TestTodo) TestTodo {
	return &testTodoController{
		testTodoUsecase: pt,
	}
}

func (c *testTodoController) Get(ctx context.Context, where *model.TestTodoWhereInput) (*model.TestTodo, error) {
	return c.testTodoUsecase.Get(ctx, where)
}

func (c *testTodoController) List(ctx context.Context) ([]*model.TestTodo, error) {
	return c.testTodoUsecase.List(ctx)
}

func (c *testTodoController) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TestTodoWhereInput, requestedFields []string) (*model.TestTodoConnection, error) {
	return c.testTodoUsecase.ListWithPagination(ctx, after, first, before, last, where, requestedFields)
}

func (c *testTodoController) Create(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error) {
	return c.testTodoUsecase.Create(ctx, input)
}

func (c *testTodoController) Update(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error) {
	return c.testTodoUsecase.Update(ctx, input)
}

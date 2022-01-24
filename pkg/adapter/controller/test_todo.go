package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TestTodo is an interface of controller
type TestTodo interface {
	Get(ctx context.Context, id *model.ID) (*model.TestTodo, error)
	List(ctx context.Context) ([]*model.TestTodo, error)
	Create(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error)
	Update(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error)
}

type testTodoController struct {
	testTodoUsecase usecase.TestTodo
}

// NewTestTodoController generates test user controller
func NewTestTodoController(tu usecase.TestTodo) TestTodo {
	return &testTodoController{
		testTodoUsecase: tu,
	}
}

func (c *testTodoController) Get(ctx context.Context, id *model.ID) (*model.TestTodo, error) {
	return c.testTodoUsecase.Get(ctx, id)
}

func (c *testTodoController) List(ctx context.Context) ([]*model.TestTodo, error) {
	return c.testTodoUsecase.List(ctx)
}

func (c *testTodoController) Create(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error) {
	return c.testTodoUsecase.Create(ctx, input)
}

func (c *testTodoController) Update(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error) {
	return c.testTodoUsecase.Update(ctx, input)
}

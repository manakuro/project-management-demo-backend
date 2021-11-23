package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TestTodo is an interface of controller
type TestTodo interface {
	Get(ctx context.Context, id *string) (*model.TestTodo, error)
	Create(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error)
	Update(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error)
}

type testTodo struct {
	testTodoUsecase usecase.TestTodo
}

// NewTestTodoController generates test user controller
func NewTestTodoController(tu usecase.TestTodo) TestTodo {
	return &testTodo{
		testTodoUsecase: tu,
	}
}

func (c *testTodo) Get(ctx context.Context, id *string) (*model.TestTodo, error) {
	return c.testTodoUsecase.Get(id)
}

func (c *testTodo) Create(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error) {
	return c.testTodoUsecase.Create(input)
}

func (c *testTodo) Update(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error) {
	return c.testTodoUsecase.Update(input)
}

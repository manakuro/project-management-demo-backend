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

type testTodo struct {
	testTodoUsecase usecase.TestTodo
}

// NewTestTodoController generates test user controller
func NewTestTodoController(tu usecase.TestTodo) TestTodo {
	return &testTodo{
		testTodoUsecase: tu,
	}
}

func (t *testTodo) Get(ctx context.Context, id *model.ID) (*model.TestTodo, error) {
	return t.testTodoUsecase.Get(ctx, id)
}

func (t *testTodo) List(ctx context.Context) ([]*model.TestTodo, error) {
	return t.testTodoUsecase.List(ctx)
}

func (t *testTodo) Create(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error) {
	return t.testTodoUsecase.Create(ctx, input)
}

func (t *testTodo) Update(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error) {
	return t.testTodoUsecase.Update(ctx, input)
}

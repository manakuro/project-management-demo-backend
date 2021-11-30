package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type testTodo struct {
	testTodoRepository repository.TestTodo
}

// TestTodo is an interface of test user
type TestTodo interface {
	Get(ctx context.Context, id *model.ID) (*model.TestTodo, error)
	List(ctx context.Context) ([]*model.TestTodo, error)
	Create(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error)
	Update(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error)
}

// NewTestTodoUsecase generates test user repository
func NewTestTodoUsecase(r repository.TestTodo) TestTodo {
	return &testTodo{testTodoRepository: r}
}

func (t *testTodo) Get(ctx context.Context, id *model.ID) (*model.TestTodo, error) {
	return t.testTodoRepository.Get(ctx, id)
}

func (t *testTodo) List(ctx context.Context) ([]*model.TestTodo, error) {
	return t.testTodoRepository.List(ctx)
}

func (t *testTodo) Create(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error) {
	return t.testTodoRepository.Create(ctx, input)
}

func (t *testTodo) Update(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error) {
	return t.testTodoRepository.Update(ctx, input)
}

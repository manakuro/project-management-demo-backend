package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type testTodoUsecase struct {
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
	return &testTodoUsecase{testTodoRepository: r}
}

func (u *testTodoUsecase) Get(ctx context.Context, id *model.ID) (*model.TestTodo, error) {
	return u.testTodoRepository.Get(ctx, id)
}

func (u *testTodoUsecase) List(ctx context.Context) ([]*model.TestTodo, error) {
	return u.testTodoRepository.List(ctx)
}

func (u *testTodoUsecase) Create(ctx context.Context, input model.CreateTestTodoInput) (*model.TestTodo, error) {
	return u.testTodoRepository.Create(ctx, input)
}

func (u *testTodoUsecase) Update(ctx context.Context, input model.UpdateTestTodoInput) (*model.TestTodo, error) {
	return u.testTodoRepository.Update(ctx, input)
}

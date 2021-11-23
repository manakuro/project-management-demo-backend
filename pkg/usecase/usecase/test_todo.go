package usecase

import (
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type testTodo struct {
	testTodoRepository repository.TestTodo
}

// TestTodo is an interface of test user
type TestTodo interface {
	Get(id *string) (*model.TestTodo, error)
	Create(input model.CreateTestTodoInput) (*model.TestTodo, error)
	Update(input model.UpdateTestTodoInput) (*model.TestTodo, error)
}

// NewTestTodoUsecase generates test user repository
func NewTestTodoUsecase(r repository.TestTodo) TestTodo {
	return &testTodo{testTodoRepository: r}
}

func (r *testTodo) Get(id *string) (*model.TestTodo, error) {
	return r.testTodoRepository.FindBy(id)
}

func (r *testTodo) Create(input model.CreateTestTodoInput) (*model.TestTodo, error) {
	return r.testTodoRepository.Create(input)
}

func (r *testTodo) Update(input model.UpdateTestTodoInput) (*model.TestTodo, error) {
	return r.testTodoRepository.Update(input)
}

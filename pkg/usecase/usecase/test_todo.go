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

func (t *testTodo) Get(id *string) (*model.TestTodo, error) {
	return t.testTodoRepository.FindBy(id)
}

func (t *testTodo) Create(input model.CreateTestTodoInput) (*model.TestTodo, error) {
	return t.testTodoRepository.Create(input)
}

func (t *testTodo) Update(input model.UpdateTestTodoInput) (*model.TestTodo, error) {
	return t.testTodoRepository.Update(input)
}

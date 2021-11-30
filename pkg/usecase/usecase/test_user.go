package usecase

import (
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type testUser struct {
	testUserRepository repository.TestUser
}

// TestUser is an interface of test user
type TestUser interface {
	Get(id *model.ID, age *int) (*model.TestUser, error)
	List() ([]*model.TestUser, error)
	ListWithPagination(after *model.Cursor, first *int, before *model.Cursor, last *int) (*model.TestUserConnection, error)
	Create(input model.CreateTestUserInput) (*model.TestUser, error)
	Update(input model.UpdateTestUserInput) (*model.TestUser, error)
}

// NewTestUserUsecase generates test user repository
func NewTestUserUsecase(r repository.TestUser) TestUser {
	return &testUser{testUserRepository: r}
}

func (t *testUser) Get(id *model.ID, age *int) (*model.TestUser, error) {
	return t.testUserRepository.Get(id, age)
}

func (t *testUser) List() ([]*model.TestUser, error) {
	return t.testUserRepository.List()
}

func (t *testUser) ListWithPagination(after *model.Cursor, first *int, before *model.Cursor, last *int) (*model.TestUserConnection, error) {
	return t.testUserRepository.ListWithPagination(after, first, before, last)
}

func (t *testUser) Create(input model.CreateTestUserInput) (*model.TestUser, error) {
	return t.testUserRepository.Create(input)
}

func (t *testUser) Update(input model.UpdateTestUserInput) (*model.TestUser, error) {
	return t.testUserRepository.Update(input)
}

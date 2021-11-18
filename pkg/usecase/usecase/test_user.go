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
	Get(id *string, age *int) (*model.TestUser, error)
	Create(input model.CreateTestUserInput) (*model.TestUser, error)
	Update(input model.UpdateTestUserInput) (*model.TestUser, error)
}

// NewTestUserUsecase generates test user repository
func NewTestUserUsecase(r repository.TestUser) TestUser {
	return &testUser{testUserRepository: r}
}

func (r *testUser) Get(id *string, age *int) (*model.TestUser, error) {
	return r.testUserRepository.FindBy(id, age)
}

func (r *testUser) Create(input model.CreateTestUserInput) (*model.TestUser, error) {
	return r.testUserRepository.Create(input)
}

func (r *testUser) Update(input model.UpdateTestUserInput) (*model.TestUser, error) {
	return r.testUserRepository.Update(input)
}

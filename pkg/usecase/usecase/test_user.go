package usecase

import (
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type testUserUsecase struct {
	testUserRepository repository.TestUser
}

// TestUserUsecase is an interface of test user
type TestUserUsecase interface {
	Get() (*model.TestUser, error)
	Create(input model.CreateTestUserInput) (*model.TestUser, error)
	Update(input model.UpdateTestUserInput) (*model.TestUser, error)
}

// NewTestUserUsecase generates test user repository
func NewTestUserUsecase(r repository.TestUser) TestUserUsecase {
	return &testUserUsecase{testUserRepository: r}
}

func (r *testUserUsecase) Get() (*model.TestUser, error) {
	return r.testUserRepository.Find()
}

func (r *testUserUsecase) Create(input model.CreateTestUserInput) (*model.TestUser, error) {
	return r.testUserRepository.Create(input)
}

func (r *testUserUsecase) Update(input model.UpdateTestUserInput) (*model.TestUser, error) {
	return r.testUserRepository.Update(input)
}

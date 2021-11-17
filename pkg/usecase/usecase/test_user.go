package usecase

import (
	"errors"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type testUserUsecase struct {
	testUserRepository repository.TestUser
}

// TestUserUsecase is an interface of test user
type TestUserUsecase interface {
	Get() (*model.TestUser, error)
}

// NewTestUserUsecase generates test user repository
func NewTestUserUsecase(r repository.TestUser) TestUserUsecase {
	return &testUserUsecase{testUserRepository: r}
}

func (r *testUserUsecase) Get() (*model.TestUser, error) {
	u, err := r.testUserRepository.Find()
	if !errors.Is(err, nil) {
		return nil, err
	}

	return u, nil
}

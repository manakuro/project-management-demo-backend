package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TestUserController is an interface of controller
type TestUserController interface {
	Get(ctx context.Context) (*model.TestUser, error)
}

type testUserController struct {
	testUserUsecase usecase.TestUserUsecase
}

// NewTestUserController generates test user controller
func NewTestUserController(tu usecase.TestUserUsecase) TestUserController {
	return &testUserController{
		testUserUsecase: tu,
	}
}

func (c *testUserController) Get(ctx context.Context) (*model.TestUser, error) {
	return c.testUserUsecase.Get()
}

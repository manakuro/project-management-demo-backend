package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TestUser is an interface of controller
type TestUser interface {
	Get(ctx context.Context, id *string, age *int) (*model.TestUser, error)
	Create(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error)
	Update(ctx context.Context, input model.UpdateTestUserInput) (*model.TestUser, error)
}

type testUser struct {
	testUserUsecase usecase.TestUser
}

// NewTestUserController generates test user controller
func NewTestUserController(tu usecase.TestUser) TestUser {
	return &testUser{
		testUserUsecase: tu,
	}
}

func (c *testUser) Get(ctx context.Context, id *string, age *int) (*model.TestUser, error) {
	return c.testUserUsecase.Get(id, age)
}

func (c *testUser) Create(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error) {
	return c.testUserUsecase.Create(input)
}

func (c *testUser) Update(ctx context.Context, input model.UpdateTestUserInput) (*model.TestUser, error) {
	return c.testUserUsecase.Update(input)
}

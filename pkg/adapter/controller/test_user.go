package controller

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/usecase"
)

// TestUser is an interface of controller
type TestUser interface {
	Get(ctx context.Context, id *model.ID, age *int) (*model.TestUser, error)
	List(ctx context.Context) ([]*model.TestUser, error)
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

func (t *testUser) Get(ctx context.Context, id *model.ID, age *int) (*model.TestUser, error) {
	return t.testUserUsecase.Get(id, age)
}

func (t *testUser) List(ctx context.Context) ([]*model.TestUser, error) {
	return t.testUserUsecase.List()
}

func (t *testUser) Create(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error) {
	return t.testUserUsecase.Create(input)
}

func (t *testUser) Update(ctx context.Context, input model.UpdateTestUserInput) (*model.TestUser, error) {
	return t.testUserUsecase.Update(input)
}

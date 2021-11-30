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
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int) (*model.TestUserConnection, error)
	Create(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error)
	CreateWithTodo(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error)
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
	return t.testUserUsecase.Get(ctx, id, age)
}

func (t *testUser) List(ctx context.Context) ([]*model.TestUser, error) {
	return t.testUserUsecase.List(ctx)
}

func (t *testUser) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int) (*model.TestUserConnection, error) {
	return t.testUserUsecase.ListWithPagination(ctx, after, first, before, last)
}

func (t *testUser) Create(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error) {
	return t.testUserUsecase.Create(ctx, input)
}

func (t *testUser) CreateWithTodo(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error) {
	return t.testUserUsecase.CreateWithTodo(ctx, input)
}

func (t *testUser) Update(ctx context.Context, input model.UpdateTestUserInput) (*model.TestUser, error) {
	return t.testUserUsecase.Update(ctx, input)
}

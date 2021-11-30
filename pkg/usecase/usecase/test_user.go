package usecase

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/usecase/repository"
)

type testUser struct {
	testUserRepository repository.TestUser
}

// TestUser is an interface of test user
type TestUser interface {
	Get(ctx context.Context, id *model.ID, age *int) (*model.TestUser, error)
	List(ctx context.Context) ([]*model.TestUser, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int) (*model.TestUserConnection, error)
	Create(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error)
	CreateWithTodo(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error)
	Update(ctx context.Context, input model.UpdateTestUserInput) (*model.TestUser, error)
}

// NewTestUserUsecase generates test user repository
func NewTestUserUsecase(r repository.TestUser) TestUser {
	return &testUser{testUserRepository: r}
}

func (t *testUser) Get(ctx context.Context, id *model.ID, age *int) (*model.TestUser, error) {
	return t.testUserRepository.Get(ctx, id, age)
}

func (t *testUser) List(ctx context.Context) ([]*model.TestUser, error) {
	return t.testUserRepository.List(ctx)
}

func (t *testUser) ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int) (*model.TestUserConnection, error) {
	return t.testUserRepository.ListWithPagination(ctx, after, first, before, last)
}

func (t *testUser) Create(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error) {
	return t.testUserRepository.Create(ctx, input)
}

func (t *testUser) CreateWithTodo(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error) {
	return t.testUserRepository.CreateWithTodo(ctx, input)
}

func (t *testUser) Update(ctx context.Context, input model.UpdateTestUserInput) (*model.TestUser, error) {
	return t.testUserRepository.Update(ctx, input)
}

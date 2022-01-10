package repository

import (
	"context"
	"project-management-demo-backend/pkg/entity/model"
)

// TestUser is interface of repository
type TestUser interface {
	Get(ctx context.Context, id model.ID, age *int) (*model.TestUser, error)
	List(ctx context.Context) ([]*model.TestUser, error)
	ListWithPagination(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TestUserWhereInput, requestedFields []string) (*model.TestUserConnection, error)
	Create(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error)
	CreateWithTodo(ctx context.Context, input model.CreateTestUserInput) (*model.TestUser, error)
	Update(ctx context.Context, input model.UpdateTestUserInput) (*model.TestUser, error)
}

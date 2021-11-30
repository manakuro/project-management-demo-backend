package repository

import "project-management-demo-backend/pkg/entity/model"

// TestUser is interface of repository
type TestUser interface {
	Get(id *model.ID, age *int) (*model.TestUser, error)
	List() ([]*model.TestUser, error)
	ListWithPagination(after *model.Cursor, first *int, before *model.Cursor, last *int) (*model.TestUserConnection, error)
	Create(input model.CreateTestUserInput) (*model.TestUser, error)
	Update(input model.UpdateTestUserInput) (*model.TestUser, error)
}

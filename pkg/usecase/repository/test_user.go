package repository

import "project-management-demo-backend/pkg/entity/model"

// TestUser is interface of repository
type TestUser interface {
	Get(id *string, age *int) (*model.TestUser, error)
	List() ([]*model.TestUser, error)
	Create(input model.CreateTestUserInput) (*model.TestUser, error)
	Update(input model.UpdateTestUserInput) (*model.TestUser, error)
}

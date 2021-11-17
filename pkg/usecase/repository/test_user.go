package repository

import "project-management-demo-backend/pkg/entity/model"

// TestUser is interface of repository
type TestUser interface {
	Find() (*model.TestUser, error)
}

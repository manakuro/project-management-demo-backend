package model

import (
	"project-management-demo-backend/ent"
)

// TestUser is the model entity for the TestUser schema.
type TestUser = ent.TestUser

// TestUserConnection is the connection containing edges to TestUser.
type TestUserConnection = ent.TestUserConnection

// CreateTestUserInput represents a mutation input for creating test users.
type CreateTestUserInput = ent.CreateTestUserInput

// UpdateTestUserInput represents a mutation input for updating test users.
type UpdateTestUserInput = ent.UpdateTestUserInput

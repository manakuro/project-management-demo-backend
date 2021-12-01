package model

import (
	"project-management-demo-backend/ent"
)

// TestUser is the model entity for the TestUser schema.
type TestUser = ent.TestUser

// TestUserConnection is the connection containing edges to TestUser.
type TestUserConnection = ent.TestUserConnection

// TestUserWhereInput represents a where input for filtering TestUser queries.
type TestUserWhereInput = ent.TestUserWhereInput

// CreateTestUserInput represents a mutation input for creating test users.
type CreateTestUserInput = ent.CreateTestUserInput

// UpdateTestUserInput represents a mutation input for updating test users.
type UpdateTestUserInput = ent.UpdateTestUserInput

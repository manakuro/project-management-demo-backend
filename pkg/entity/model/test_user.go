package model

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/util/datetime"
)

// TestUser is the model entity for the TestUser schema.
type TestUser struct {
	*ent.TestUser
}

// CreateTestUserInput represents a mutation input for creating test users.
type CreateTestUserInput struct {
	ent.CreateTestUserInput
}

// UpdateTestUserInput represents a mutation input for updating test users.
type UpdateTestUserInput struct {
	ID string
	ent.UpdateTestUserInput
}

// FormattedCreatedAt returns formatted created_at
func (t *TestUser) FormattedCreatedAt() string {
	return datetime.FormatDate(t.CreatedAt)
}

// FormattedUpdatedAt returns formatted updated_at
func (t *TestUser) FormattedUpdatedAt() string {
	return datetime.FormatDate(t.UpdatedAt)
}

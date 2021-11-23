package model

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/util/datetime"
)

// TestTodo is the model entity for the TestTodo schema
type TestTodo struct {
	*ent.TestTodo
}

// CreateTestTodoInput represents a mutation input for creating test users.
type CreateTestTodoInput struct {
	ent.CreateTestTodoInput
}

// UpdateTestTodoInput represents a mutation input for updating test users.
type UpdateTestTodoInput struct {
	ID string
	ent.UpdateTestTodoInput
}

// FormattedCreatedAt returns formatted created_at
func (t *TestTodo) FormattedCreatedAt() string {
	return datetime.FormatDate(t.CreatedAt)
}

// FormattedUpdatedAt returns formatted updated_at
func (t *TestTodo) FormattedUpdatedAt() string {
	return datetime.FormatDate(t.UpdatedAt)
}

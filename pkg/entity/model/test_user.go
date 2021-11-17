package model

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/util/datetime"
)

// TestUser is the model entity for the TestUser schema.
type TestUser struct {
	*ent.TestUser
}

// FormattedCreatedAt returns formatted created_at
func (t *TestUser) FormattedCreatedAt() string {
	return datetime.FormatDate(t.CreatedAt)
}

// FormattedUpdatedAt returns formatted updated_at
func (t *TestUser) FormattedUpdatedAt() string {
	return datetime.FormatDate(t.UpdatedAt)
}

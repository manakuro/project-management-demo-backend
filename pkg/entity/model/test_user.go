package model

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/utils/datetime"
)

// TestUser is the model entity for the TestUser schema.
type TestUser struct {
	*ent.TestUser
}

// FormattedCreatedAt returns formatted created_at
func (t *TestUser) FormattedCreatedAt() string {
	return datetime.FormatDate(t.CreatedAt)
}

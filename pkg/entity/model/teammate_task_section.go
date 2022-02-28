package model

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
)

// TeammateTaskSection is the model entity for the TeammateTaskSection schema.
type TeammateTaskSection = ent.TeammateTaskSection

// TeammateTaskSectionConnection is the connection containing edges to TeammateTaskSection.
type TeammateTaskSectionConnection = ent.TeammateTaskSectionConnection

// CreateTeammateTaskSectionInput represents a mutation input.
type CreateTeammateTaskSectionInput = ent.CreateTeammateTaskSectionInput

// TeammateTaskSectionWhereInput represents a where input.
type TeammateTaskSectionWhereInput = ent.TeammateTaskSectionWhereInput

// UpdateTeammateTaskSectionInput represents a mutation input.
type UpdateTeammateTaskSectionInput = ent.UpdateTeammateTaskSectionInput

// DeleteTeammateTaskSectionInput represents a mutation input.
type DeleteTeammateTaskSectionInput struct {
	ID ulid.ID
}

// DeleteTeammateTaskSectionAndKeepTasksInput represents a mutation input.
type DeleteTeammateTaskSectionAndKeepTasksInput struct {
	ID        ulid.ID
	RequestID string
}

// DeleteTeammateTaskSectionAndKeepTasksPayload represents a mutation payload.
type DeleteTeammateTaskSectionAndKeepTasksPayload struct {
	TeammateTaskSection *TeammateTaskSection
	TeammateTaskIDs     []ulid.ID
}

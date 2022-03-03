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
	ID          ulid.ID
	WorkspaceID ulid.ID
	RequestID   string
}

// DeleteTeammateTaskSectionAndKeepTasksInput represents a mutation input.
type DeleteTeammateTaskSectionAndKeepTasksInput struct {
	ID          ulid.ID
	WorkspaceID ulid.ID
	RequestID   string
}

// DeleteTeammateTaskSectionAndKeepTasksPayload represents a mutation payload.
type DeleteTeammateTaskSectionAndKeepTasksPayload struct {
	TeammateTaskSection     *TeammateTaskSection
	KeptTeammateTaskSection *TeammateTaskSection
	TeammateTaskIDs         []ulid.ID
}

// DeleteTeammateTaskSectionAndDeleteTasksInput represents a mutation input.
type DeleteTeammateTaskSectionAndDeleteTasksInput struct {
	ID          ulid.ID
	WorkspaceID ulid.ID
	RequestID   string
}

// DeleteTeammateTaskSectionAndDeleteTasksPayload represents a mutation payload.
type DeleteTeammateTaskSectionAndDeleteTasksPayload struct {
	TeammateTaskSection *TeammateTaskSection
	TeammateTaskIDs     []ulid.ID
}

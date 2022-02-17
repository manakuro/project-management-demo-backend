package model

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
)

// TaskCollaborator is the model entity for the TaskCollaborator schema.
type TaskCollaborator = ent.TaskCollaborator

// TaskCollaboratorConnection is the connection containing edges to TaskCollaborator.
type TaskCollaboratorConnection = ent.TaskCollaboratorConnection

// CreateTaskCollaboratorInput represents a mutation input.
type CreateTaskCollaboratorInput = ent.CreateTaskCollaboratorInput

// TaskCollaboratorWhereInput represents a where input.
type TaskCollaboratorWhereInput = ent.TaskCollaboratorWhereInput

// UpdateTaskCollaboratorInput represents a mutation input.
type UpdateTaskCollaboratorInput = ent.UpdateTaskCollaboratorInput

// DeleteTaskCollaboratorInput represents a mutation input.
type DeleteTaskCollaboratorInput struct {
	ID        ulid.ID
	RequestID string
}

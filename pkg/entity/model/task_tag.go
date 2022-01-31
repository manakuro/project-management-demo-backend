package model

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
)

// TaskTag is the model entity for the TaskTag schema.
type TaskTag = ent.TaskTag

// TaskTagConnection is the connection containing edges to TaskTag.
type TaskTagConnection = ent.TaskTagConnection

// CreateTaskTagInput represents a mutation input.
type CreateTaskTagInput = ent.CreateTaskTagInput

// TaskTagWhereInput represents a where input.
type TaskTagWhereInput = ent.TaskTagWhereInput

// UpdateTaskTagInput represents a mutation input.
type UpdateTaskTagInput = ent.UpdateTaskTagInput

// DeleteTaskTagInput represents a mutation input.
type DeleteTaskTagInput struct {
	ID ulid.ID
}

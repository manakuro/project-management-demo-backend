package model

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
)

// TeammateTask is the model entity for the TeammateTask schema.
type TeammateTask = ent.TeammateTask

// TeammateTaskConnection is the connection containing edges to TeammateTask.
type TeammateTaskConnection = ent.TeammateTaskConnection

// CreateTeammateTaskInput represents a mutation input.
type CreateTeammateTaskInput = ent.CreateTeammateTaskInput

// TeammateTaskWhereInput represents a where input.
type TeammateTaskWhereInput = ent.TeammateTaskWhereInput

// UpdateTeammateTaskInput represents a mutation input.
type UpdateTeammateTaskInput = ent.UpdateTeammateTaskInput

// DeleteTeammateTaskInput represents a mutation input.
type DeleteTeammateTaskInput struct {
	ID ulid.ID
}

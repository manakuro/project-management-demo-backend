package model

import (
	"project-management-demo-backend/ent"
)

// Task is the model entity for the Task schema
type Task = ent.Task

// TaskConnection is the connection containing edges to Task.
type TaskConnection = ent.TaskConnection

// CreateTaskInput represents a mutation input for creating favorite workspace.
type CreateTaskInput = ent.CreateTaskInput

// TaskWhereInput represents a where input for filtering Task queries.
type TaskWhereInput = ent.TaskWhereInput

// UpdateTaskInput represents a mutation input for updating favorite workspace.
type UpdateTaskInput = ent.UpdateTaskInput

package model

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
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

// DeleteTaskInput represents a mutation input.
type DeleteTaskInput struct {
	TaskID      ulid.ID
	WorkspaceID ulid.ID
	RequestID   string
}

// DeleteTaskPayload represents a mutation result.
type DeleteTaskPayload struct {
	TeammateTask *TeammateTask
	ProjectTask  *ProjectTask
	DeletedTasks []*DeletedTask
}

// UndeleteTaskInput represents a mutation input.
type UndeleteTaskInput struct {
	TaskID      ulid.ID
	WorkspaceID ulid.ID
	RequestID   string
}

// UndeleteTaskPayload represents a mutation result.
type UndeleteTaskPayload struct {
	TeammateTask *TeammateTask
	ProjectTask  *ProjectTask
	DeletedTasks []*DeletedTask
}

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

// DeleteTaskInput represents a mutation input.
type DeleteTaskInput struct {
	TaskID      ID
	WorkspaceID ID
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
	TaskID      ID
	WorkspaceID ID
	RequestID   string
}

// UndeleteTaskPayload represents a mutation result.
type UndeleteTaskPayload struct {
	TeammateTask *TeammateTask
	ProjectTask  *ProjectTask
	DeletedTasks []*DeletedTask
}

// DeleteAllTaskInput represents a mutation input.
type DeleteAllTaskInput struct {
	TaskIDs     []ID
	WorkspaceID ID
	RequestID   string
}

// DeleteAllTaskPayload represents a mutation result.
type DeleteAllTaskPayload struct {
	TeammateTasks []*TeammateTask
	ProjectTasks  []*ProjectTask
	DeletedTasks  []*DeletedTask
}

// UndeleteAllTaskInput represents a mutation input.
type UndeleteAllTaskInput struct {
	TaskIDs               []ID
	WorkspaceID           ID
	TeammateTaskSectionID *ID
	ProjectTaskSectionID  *ID
	RequestID             string
}

// UndeleteAllTaskPayload represents a mutation result.
type UndeleteAllTaskPayload struct {
	TeammateTasks []*TeammateTask
	ProjectTasks  []*ProjectTask
	DeletedTasks  []*DeletedTask
}

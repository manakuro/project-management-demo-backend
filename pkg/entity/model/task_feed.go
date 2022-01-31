package model

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
)

// TaskFeed is the model entity for the TaskFeed schema.
type TaskFeed = ent.TaskFeed

// TaskFeedConnection is the connection containing edges to TaskFeed.
type TaskFeedConnection = ent.TaskFeedConnection

// CreateTaskFeedInput represents a mutation input.
type CreateTaskFeedInput = ent.CreateTaskFeedInput

// TaskFeedWhereInput represents a where input.
type TaskFeedWhereInput = ent.TaskFeedWhereInput

// UpdateTaskFeedInput represents a mutation input.
type UpdateTaskFeedInput = ent.UpdateTaskFeedInput

// DeleteTaskFeedInput represents a mutation input.
type DeleteTaskFeedInput struct {
	ID ulid.ID
}

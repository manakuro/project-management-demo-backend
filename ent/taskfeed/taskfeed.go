// Code generated by entc, DO NOT EDIT.

package taskfeed

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the taskfeed type in the database.
	Label = "task_feed"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTaskID holds the string denoting the task_id field in the database.
	FieldTaskID = "task_id"
	// FieldTeammateID holds the string denoting the teammate_id field in the database.
	FieldTeammateID = "teammate_id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIsFirst holds the string denoting the is_first field in the database.
	FieldIsFirst = "is_first"
	// FieldIsPinned holds the string denoting the is_pinned field in the database.
	FieldIsPinned = "is_pinned"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "task"
	// EdgeTeammate holds the string denoting the teammate edge name in mutations.
	EdgeTeammate = "teammate"
	// EdgeTaskFeedLikes holds the string denoting the taskfeedlikes edge name in mutations.
	EdgeTaskFeedLikes = "taskFeedLikes"
	// EdgeTaskFiles holds the string denoting the taskfiles edge name in mutations.
	EdgeTaskFiles = "taskFiles"
	// Table holds the table name of the taskfeed in the database.
	Table = "task_feeds"
	// TaskTable is the table that holds the task relation/edge.
	TaskTable = "task_feeds"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "tasks"
	// TaskColumn is the table column denoting the task relation/edge.
	TaskColumn = "task_id"
	// TeammateTable is the table that holds the teammate relation/edge.
	TeammateTable = "task_feeds"
	// TeammateInverseTable is the table name for the Teammate entity.
	// It exists in this package in order to avoid circular dependency with the "teammate" package.
	TeammateInverseTable = "teammates"
	// TeammateColumn is the table column denoting the teammate relation/edge.
	TeammateColumn = "teammate_id"
	// TaskFeedLikesTable is the table that holds the taskFeedLikes relation/edge.
	TaskFeedLikesTable = "task_feed_likes"
	// TaskFeedLikesInverseTable is the table name for the TaskFeedLike entity.
	// It exists in this package in order to avoid circular dependency with the "taskfeedlike" package.
	TaskFeedLikesInverseTable = "task_feed_likes"
	// TaskFeedLikesColumn is the table column denoting the taskFeedLikes relation/edge.
	TaskFeedLikesColumn = "task_feed_id"
	// TaskFilesTable is the table that holds the taskFiles relation/edge.
	TaskFilesTable = "task_files"
	// TaskFilesInverseTable is the table name for the TaskFile entity.
	// It exists in this package in order to avoid circular dependency with the "taskfile" package.
	TaskFilesInverseTable = "task_files"
	// TaskFilesColumn is the table column denoting the taskFiles relation/edge.
	TaskFilesColumn = "task_feed_id"
)

// Columns holds all SQL columns for taskfeed fields.
var Columns = []string{
	FieldID,
	FieldTaskID,
	FieldTeammateID,
	FieldDescription,
	FieldIsFirst,
	FieldIsPinned,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsFirst holds the default value on creation for the "is_first" field.
	DefaultIsFirst bool
	// DefaultIsPinned holds the default value on creation for the "is_pinned" field.
	DefaultIsPinned bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)

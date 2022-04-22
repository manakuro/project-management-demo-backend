// Code generated by entc, DO NOT EDIT.

package deletedteammatetask

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the deletedteammatetask type in the database.
	Label = "deleted_teammate_task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTeammateID holds the string denoting the teammate_id field in the database.
	FieldTeammateID = "teammate_id"
	// FieldTaskID holds the string denoting the task_id field in the database.
	FieldTaskID = "task_id"
	// FieldTeammateTaskSectionID holds the string denoting the teammate_task_section_id field in the database.
	FieldTeammateTaskSectionID = "teammate_task_section_id"
	// FieldWorkspaceID holds the string denoting the workspace_id field in the database.
	FieldWorkspaceID = "workspace_id"
	// FieldTeammateTaskID holds the string denoting the teammate_task_id field in the database.
	FieldTeammateTaskID = "teammate_task_id"
	// FieldTeammateTaskCreatedAt holds the string denoting the teammate_task_created_at field in the database.
	FieldTeammateTaskCreatedAt = "teammate_task_created_at"
	// FieldTeammateTaskUpdatedAt holds the string denoting the teammate_task_updated_at field in the database.
	FieldTeammateTaskUpdatedAt = "teammate_task_updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTeammate holds the string denoting the teammate edge name in mutations.
	EdgeTeammate = "teammate"
	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "task"
	// EdgeWorkspace holds the string denoting the workspace edge name in mutations.
	EdgeWorkspace = "workspace"
	// Table holds the table name of the deletedteammatetask in the database.
	Table = "deleted_teammate_tasks"
	// TeammateTable is the table that holds the teammate relation/edge.
	TeammateTable = "deleted_teammate_tasks"
	// TeammateInverseTable is the table name for the Teammate entity.
	// It exists in this package in order to avoid circular dependency with the "teammate" package.
	TeammateInverseTable = "teammates"
	// TeammateColumn is the table column denoting the teammate relation/edge.
	TeammateColumn = "teammate_id"
	// TaskTable is the table that holds the task relation/edge.
	TaskTable = "deleted_teammate_tasks"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "tasks"
	// TaskColumn is the table column denoting the task relation/edge.
	TaskColumn = "task_id"
	// WorkspaceTable is the table that holds the workspace relation/edge.
	WorkspaceTable = "deleted_teammate_tasks"
	// WorkspaceInverseTable is the table name for the Workspace entity.
	// It exists in this package in order to avoid circular dependency with the "workspace" package.
	WorkspaceInverseTable = "workspaces"
	// WorkspaceColumn is the table column denoting the workspace relation/edge.
	WorkspaceColumn = "workspace_id"
)

// Columns holds all SQL columns for deletedteammatetask fields.
var Columns = []string{
	FieldID,
	FieldTeammateID,
	FieldTaskID,
	FieldTeammateTaskSectionID,
	FieldWorkspaceID,
	FieldTeammateTaskID,
	FieldTeammateTaskCreatedAt,
	FieldTeammateTaskUpdatedAt,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)
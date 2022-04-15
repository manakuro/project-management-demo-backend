// Code generated by entc, DO NOT EDIT.

package archivedtaskactivity

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the archivedtaskactivity type in the database.
	Label = "archived_task_activity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldActivityTypeID holds the string denoting the activity_type_id field in the database.
	FieldActivityTypeID = "activity_type_id"
	// FieldTeammateID holds the string denoting the teammate_id field in the database.
	FieldTeammateID = "teammate_id"
	// FieldWorkspaceID holds the string denoting the workspace_id field in the database.
	FieldWorkspaceID = "workspace_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTeammate holds the string denoting the teammate edge name in mutations.
	EdgeTeammate = "teammate"
	// EdgeActivityType holds the string denoting the activitytype edge name in mutations.
	EdgeActivityType = "activityType"
	// EdgeWorkspace holds the string denoting the workspace edge name in mutations.
	EdgeWorkspace = "workspace"
	// EdgeArchivedTaskActivityTasks holds the string denoting the archivedtaskactivitytasks edge name in mutations.
	EdgeArchivedTaskActivityTasks = "archivedTaskActivityTasks"
	// Table holds the table name of the archivedtaskactivity in the database.
	Table = "archived_task_activities"
	// TeammateTable is the table that holds the teammate relation/edge.
	TeammateTable = "archived_task_activities"
	// TeammateInverseTable is the table name for the Teammate entity.
	// It exists in this package in order to avoid circular dependency with the "teammate" package.
	TeammateInverseTable = "teammates"
	// TeammateColumn is the table column denoting the teammate relation/edge.
	TeammateColumn = "teammate_id"
	// ActivityTypeTable is the table that holds the activityType relation/edge.
	ActivityTypeTable = "archived_task_activities"
	// ActivityTypeInverseTable is the table name for the ActivityType entity.
	// It exists in this package in order to avoid circular dependency with the "activitytype" package.
	ActivityTypeInverseTable = "activity_types"
	// ActivityTypeColumn is the table column denoting the activityType relation/edge.
	ActivityTypeColumn = "activity_type_id"
	// WorkspaceTable is the table that holds the workspace relation/edge.
	WorkspaceTable = "archived_task_activities"
	// WorkspaceInverseTable is the table name for the Workspace entity.
	// It exists in this package in order to avoid circular dependency with the "workspace" package.
	WorkspaceInverseTable = "workspaces"
	// WorkspaceColumn is the table column denoting the workspace relation/edge.
	WorkspaceColumn = "workspace_id"
	// ArchivedTaskActivityTasksTable is the table that holds the archivedTaskActivityTasks relation/edge.
	ArchivedTaskActivityTasksTable = "archived_task_activity_tasks"
	// ArchivedTaskActivityTasksInverseTable is the table name for the ArchivedTaskActivityTask entity.
	// It exists in this package in order to avoid circular dependency with the "archivedtaskactivitytask" package.
	ArchivedTaskActivityTasksInverseTable = "archived_task_activity_tasks"
	// ArchivedTaskActivityTasksColumn is the table column denoting the archivedTaskActivityTasks relation/edge.
	ArchivedTaskActivityTasksColumn = "archived_task_activity_id"
)

// Columns holds all SQL columns for archivedtaskactivity fields.
var Columns = []string{
	FieldID,
	FieldActivityTypeID,
	FieldTeammateID,
	FieldWorkspaceID,
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

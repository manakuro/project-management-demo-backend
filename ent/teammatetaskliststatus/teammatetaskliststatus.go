// Code generated by entc, DO NOT EDIT.

package teammatetaskliststatus

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the teammatetaskliststatus type in the database.
	Label = "teammate_task_list_status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWorkspaceID holds the string denoting the workspace_id field in the database.
	FieldWorkspaceID = "workspace_id"
	// FieldTeammateID holds the string denoting the teammate_id field in the database.
	FieldTeammateID = "teammate_id"
	// FieldTaskListCompletedStatusID holds the string denoting the task_list_completed_status_id field in the database.
	FieldTaskListCompletedStatusID = "task_list_completed_status_id"
	// FieldTaskListSortStatusID holds the string denoting the task_list_sort_status_id field in the database.
	FieldTaskListSortStatusID = "task_list_sort_status_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeWorkspace holds the string denoting the workspace edge name in mutations.
	EdgeWorkspace = "workspace"
	// EdgeTeammate holds the string denoting the teammate edge name in mutations.
	EdgeTeammate = "teammate"
	// EdgeTaskListCompletedStatus holds the string denoting the tasklistcompletedstatus edge name in mutations.
	EdgeTaskListCompletedStatus = "taskListCompletedStatus"
	// EdgeTaskListSortStatus holds the string denoting the tasklistsortstatus edge name in mutations.
	EdgeTaskListSortStatus = "taskListSortStatus"
	// Table holds the table name of the teammatetaskliststatus in the database.
	Table = "teammate_task_list_status"
	// WorkspaceTable is the table that holds the workspace relation/edge.
	WorkspaceTable = "teammate_task_list_status"
	// WorkspaceInverseTable is the table name for the Workspace entity.
	// It exists in this package in order to avoid circular dependency with the "workspace" package.
	WorkspaceInverseTable = "workspaces"
	// WorkspaceColumn is the table column denoting the workspace relation/edge.
	WorkspaceColumn = "workspace_id"
	// TeammateTable is the table that holds the teammate relation/edge.
	TeammateTable = "teammate_task_list_status"
	// TeammateInverseTable is the table name for the Teammate entity.
	// It exists in this package in order to avoid circular dependency with the "teammate" package.
	TeammateInverseTable = "teammates"
	// TeammateColumn is the table column denoting the teammate relation/edge.
	TeammateColumn = "teammate_id"
	// TaskListCompletedStatusTable is the table that holds the taskListCompletedStatus relation/edge.
	TaskListCompletedStatusTable = "teammate_task_list_status"
	// TaskListCompletedStatusInverseTable is the table name for the TaskListCompletedStatus entity.
	// It exists in this package in order to avoid circular dependency with the "tasklistcompletedstatus" package.
	TaskListCompletedStatusInverseTable = "task_list_completed_status"
	// TaskListCompletedStatusColumn is the table column denoting the taskListCompletedStatus relation/edge.
	TaskListCompletedStatusColumn = "task_list_completed_status_id"
	// TaskListSortStatusTable is the table that holds the taskListSortStatus relation/edge.
	TaskListSortStatusTable = "teammate_task_list_status"
	// TaskListSortStatusInverseTable is the table name for the TaskListSortStatus entity.
	// It exists in this package in order to avoid circular dependency with the "tasklistsortstatus" package.
	TaskListSortStatusInverseTable = "task_list_sort_status"
	// TaskListSortStatusColumn is the table column denoting the taskListSortStatus relation/edge.
	TaskListSortStatusColumn = "task_list_sort_status_id"
)

// Columns holds all SQL columns for teammatetaskliststatus fields.
var Columns = []string{
	FieldID,
	FieldWorkspaceID,
	FieldTeammateID,
	FieldTaskListCompletedStatusID,
	FieldTaskListSortStatusID,
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

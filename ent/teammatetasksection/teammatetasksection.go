// Code generated by entc, DO NOT EDIT.

package teammatetasksection

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the teammatetasksection type in the database.
	Label = "teammate_task_section"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTeammateID holds the string denoting the teammate_id field in the database.
	FieldTeammateID = "teammate_id"
	// FieldWorkspaceID holds the string denoting the workspace_id field in the database.
	FieldWorkspaceID = "workspace_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAssigned holds the string denoting the assigned field in the database.
	FieldAssigned = "assigned"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTeammate holds the string denoting the teammate edge name in mutations.
	EdgeTeammate = "teammate"
	// EdgeWorkspace holds the string denoting the workspace edge name in mutations.
	EdgeWorkspace = "workspace"
	// Table holds the table name of the teammatetasksection in the database.
	Table = "teammate_task_sections"
	// TeammateTable is the table that holds the teammate relation/edge.
	TeammateTable = "teammate_task_sections"
	// TeammateInverseTable is the table name for the Teammate entity.
	// It exists in this package in order to avoid circular dependency with the "teammate" package.
	TeammateInverseTable = "teammates"
	// TeammateColumn is the table column denoting the teammate relation/edge.
	TeammateColumn = "teammate_id"
	// WorkspaceTable is the table that holds the workspace relation/edge.
	WorkspaceTable = "teammate_task_sections"
	// WorkspaceInverseTable is the table name for the Workspace entity.
	// It exists in this package in order to avoid circular dependency with the "workspace" package.
	WorkspaceInverseTable = "workspaces"
	// WorkspaceColumn is the table column denoting the workspace relation/edge.
	WorkspaceColumn = "workspace_id"
)

// Columns holds all SQL columns for teammatetasksection fields.
var Columns = []string{
	FieldID,
	FieldTeammateID,
	FieldWorkspaceID,
	FieldName,
	FieldAssigned,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)

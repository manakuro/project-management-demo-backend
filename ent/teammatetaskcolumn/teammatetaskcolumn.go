// Code generated by entc, DO NOT EDIT.

package teammatetaskcolumn

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the teammatetaskcolumn type in the database.
	Label = "teammate_task_column"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTeammateID holds the string denoting the teammate_id field in the database.
	FieldTeammateID = "teammate_id"
	// FieldTaskColumnID holds the string denoting the task_column_id field in the database.
	FieldTaskColumnID = "task_column_id"
	// FieldWidth holds the string denoting the width field in the database.
	FieldWidth = "width"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// FieldCustomizable holds the string denoting the customizable field in the database.
	FieldCustomizable = "customizable"
	// FieldOrder holds the string denoting the order field in the database.
	FieldOrder = "order"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTeammate holds the string denoting the teammate edge name in mutations.
	EdgeTeammate = "teammate"
	// EdgeTaskColumn holds the string denoting the task_column edge name in mutations.
	EdgeTaskColumn = "task_column"
	// Table holds the table name of the teammatetaskcolumn in the database.
	Table = "teammate_task_columns"
	// TeammateTable is the table that holds the teammate relation/edge.
	TeammateTable = "teammate_task_columns"
	// TeammateInverseTable is the table name for the Teammate entity.
	// It exists in this package in order to avoid circular dependency with the "teammate" package.
	TeammateInverseTable = "teammates"
	// TeammateColumn is the table column denoting the teammate relation/edge.
	TeammateColumn = "teammate_id"
	// TaskColumnTable is the table that holds the task_column relation/edge.
	TaskColumnTable = "teammate_task_columns"
	// TaskColumnInverseTable is the table name for the TaskColumn entity.
	// It exists in this package in order to avoid circular dependency with the "taskcolumn" package.
	TaskColumnInverseTable = "task_columns"
	// TaskColumnColumn is the table column denoting the task_column relation/edge.
	TaskColumnColumn = "task_column_id"
)

// Columns holds all SQL columns for teammatetaskcolumn fields.
var Columns = []string{
	FieldID,
	FieldTeammateID,
	FieldTaskColumnID,
	FieldWidth,
	FieldDisabled,
	FieldCustomizable,
	FieldOrder,
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
	// WidthValidator is a validator for the "width" field. It is called by the builders before save.
	WidthValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)

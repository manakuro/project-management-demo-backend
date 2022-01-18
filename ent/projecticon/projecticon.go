// Code generated by entc, DO NOT EDIT.

package projecticon

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the projecticon type in the database.
	Label = "project_icon"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIconID holds the string denoting the icon_id field in the database.
	FieldIconID = "icon_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeProjects holds the string denoting the projects edge name in mutations.
	EdgeProjects = "projects"
	// EdgeIcon holds the string denoting the icon edge name in mutations.
	EdgeIcon = "icon"
	// Table holds the table name of the projecticon in the database.
	Table = "project_icons"
	// ProjectsTable is the table that holds the projects relation/edge.
	ProjectsTable = "projects"
	// ProjectsInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectsInverseTable = "projects"
	// ProjectsColumn is the table column denoting the projects relation/edge.
	ProjectsColumn = "project_icon_id"
	// IconTable is the table that holds the icon relation/edge.
	IconTable = "project_icons"
	// IconInverseTable is the table name for the Icon entity.
	// It exists in this package in order to avoid circular dependency with the "icon" package.
	IconInverseTable = "icons"
	// IconColumn is the table column denoting the icon relation/edge.
	IconColumn = "icon_id"
)

// Columns holds all SQL columns for projecticon fields.
var Columns = []string{
	FieldID,
	FieldIconID,
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
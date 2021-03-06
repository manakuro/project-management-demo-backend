// Code generated by entc, DO NOT EDIT.

package taskfile

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"
)

const (
	// Label holds the string label denoting the taskfile type in the database.
	Label = "task_file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// FieldTaskID holds the string denoting the task_id field in the database.
	FieldTaskID = "task_id"
	// FieldTaskFeedID holds the string denoting the task_feed_id field in the database.
	FieldTaskFeedID = "task_feed_id"
	// FieldFileTypeID holds the string denoting the file_type_id field in the database.
	FieldFileTypeID = "file_type_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSrc holds the string denoting the src field in the database.
	FieldSrc = "src"
	// FieldAttached holds the string denoting the attached field in the database.
	FieldAttached = "attached"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "task"
	// EdgeTaskFeed holds the string denoting the taskfeed edge name in mutations.
	EdgeTaskFeed = "taskFeed"
	// EdgeFileType holds the string denoting the filetype edge name in mutations.
	EdgeFileType = "fileType"
	// Table holds the table name of the taskfile in the database.
	Table = "task_files"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "task_files"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_id"
	// TaskTable is the table that holds the task relation/edge.
	TaskTable = "task_files"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "tasks"
	// TaskColumn is the table column denoting the task relation/edge.
	TaskColumn = "task_id"
	// TaskFeedTable is the table that holds the taskFeed relation/edge.
	TaskFeedTable = "task_files"
	// TaskFeedInverseTable is the table name for the TaskFeed entity.
	// It exists in this package in order to avoid circular dependency with the "taskfeed" package.
	TaskFeedInverseTable = "task_feeds"
	// TaskFeedColumn is the table column denoting the taskFeed relation/edge.
	TaskFeedColumn = "task_feed_id"
	// FileTypeTable is the table that holds the fileType relation/edge.
	FileTypeTable = "task_files"
	// FileTypeInverseTable is the table name for the FileType entity.
	// It exists in this package in order to avoid circular dependency with the "filetype" package.
	FileTypeInverseTable = "file_types"
	// FileTypeColumn is the table column denoting the fileType relation/edge.
	FileTypeColumn = "file_type_id"
)

// Columns holds all SQL columns for taskfile fields.
var Columns = []string{
	FieldID,
	FieldProjectID,
	FieldTaskID,
	FieldTaskFeedID,
	FieldFileTypeID,
	FieldName,
	FieldSrc,
	FieldAttached,
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
	// SrcValidator is a validator for the "src" field. It is called by the builders before save.
	SrcValidator func(string) error
	// DefaultAttached holds the default value on creation for the "attached" field.
	DefaultAttached bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)

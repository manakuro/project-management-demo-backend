// Code generated by entc, DO NOT EDIT.

package taskcolumn

import (
	"fmt"
	"io"
	"project-management-demo-backend/ent/schema/ulid"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the taskcolumn type in the database.
	Label = "task_column"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTeammateTaskColumns holds the string denoting the teammate_task_columns edge name in mutations.
	EdgeTeammateTaskColumns = "teammate_task_columns"
	// Table holds the table name of the taskcolumn in the database.
	Table = "task_columns"
	// TeammateTaskColumnsTable is the table that holds the teammate_task_columns relation/edge.
	TeammateTaskColumnsTable = "teammate_task_columns"
	// TeammateTaskColumnsInverseTable is the table name for the TeammateTaskColumn entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetaskcolumn" package.
	TeammateTaskColumnsInverseTable = "teammate_task_columns"
	// TeammateTaskColumnsColumn is the table column denoting the teammate_task_columns relation/edge.
	TeammateTaskColumnsColumn = "task_column_id"
)

// Columns holds all SQL columns for taskcolumn fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
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

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeTaskName Type = "TASK_NAME"
	TypeAssignee Type = "ASSIGNEE"
	TypeDueDate  Type = "DUE_DATE"
	TypeProject  Type = "PROJECT"
	TypeTags     Type = "TAGS"
	TypeCustom   Type = "CUSTOM"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeTaskName, TypeAssignee, TypeDueDate, TypeProject, TypeTags, TypeCustom:
		return nil
	default:
		return fmt.Errorf("taskcolumn: invalid enum value for type field: %q", _type)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (_type Type) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(_type.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (_type *Type) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*_type = Type(str)
	if err := TypeValidator(*_type); err != nil {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}

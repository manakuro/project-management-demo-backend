// Code generated by entc, DO NOT EDIT.

package activitytype

import (
	"fmt"
	"io"
	"project-management-demo-backend/ent/schema/ulid"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the activitytype type in the database.
	Label = "activity_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTypeCode holds the string denoting the type_code field in the database.
	FieldTypeCode = "type_code"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTaskActivities holds the string denoting the taskactivities edge name in mutations.
	EdgeTaskActivities = "taskActivities"
	// Table holds the table name of the activitytype in the database.
	Table = "activity_types"
	// TaskActivitiesTable is the table that holds the taskActivities relation/edge.
	TaskActivitiesTable = "task_activities"
	// TaskActivitiesInverseTable is the table name for the TaskActivity entity.
	// It exists in this package in order to avoid circular dependency with the "taskactivity" package.
	TaskActivitiesInverseTable = "task_activities"
	// TaskActivitiesColumn is the table column denoting the taskActivities relation/edge.
	TaskActivitiesColumn = "activity_type_id"
)

// Columns holds all SQL columns for activitytype fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldTypeCode,
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

// TypeCode defines the type for the "type_code" enum field.
type TypeCode string

// TypeCode values.
const (
	TypeCodeTask      TypeCode = "TASK"
	TypeCodeWorkspace TypeCode = "WORKSPACE"
)

func (tc TypeCode) String() string {
	return string(tc)
}

// TypeCodeValidator is a validator for the "type_code" field enum values. It is called by the builders before save.
func TypeCodeValidator(tc TypeCode) error {
	switch tc {
	case TypeCodeTask, TypeCodeWorkspace:
		return nil
	default:
		return fmt.Errorf("activitytype: invalid enum value for type_code field: %q", tc)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (tc TypeCode) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(tc.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (tc *TypeCode) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*tc = TypeCode(str)
	if err := TypeCodeValidator(*tc); err != nil {
		return fmt.Errorf("%s is not a valid TypeCode", str)
	}
	return nil
}

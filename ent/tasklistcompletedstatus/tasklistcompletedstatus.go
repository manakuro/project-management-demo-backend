// Code generated by entc, DO NOT EDIT.

package tasklistcompletedstatus

import (
	"fmt"
	"io"
	"project-management-demo-backend/ent/schema/ulid"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the tasklistcompletedstatus type in the database.
	Label = "task_list_completed_status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStatusCode holds the string denoting the status_code field in the database.
	FieldStatusCode = "status_code"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTeammateTaskListStatuses holds the string denoting the teammate_task_list_statuses edge name in mutations.
	EdgeTeammateTaskListStatuses = "teammate_task_list_statuses"
	// EdgeProjectTaskListStatuses holds the string denoting the project_task_list_statuses edge name in mutations.
	EdgeProjectTaskListStatuses = "project_task_list_statuses"
	// Table holds the table name of the tasklistcompletedstatus in the database.
	Table = "task_list_completed_status"
	// TeammateTaskListStatusesTable is the table that holds the teammate_task_list_statuses relation/edge.
	TeammateTaskListStatusesTable = "teammate_task_list_status"
	// TeammateTaskListStatusesInverseTable is the table name for the TeammateTaskListStatus entity.
	// It exists in this package in order to avoid circular dependency with the "teammatetaskliststatus" package.
	TeammateTaskListStatusesInverseTable = "teammate_task_list_status"
	// TeammateTaskListStatusesColumn is the table column denoting the teammate_task_list_statuses relation/edge.
	TeammateTaskListStatusesColumn = "task_list_completed_status_id"
	// ProjectTaskListStatusesTable is the table that holds the project_task_list_statuses relation/edge.
	ProjectTaskListStatusesTable = "project_task_list_status"
	// ProjectTaskListStatusesInverseTable is the table name for the ProjectTaskListStatus entity.
	// It exists in this package in order to avoid circular dependency with the "projecttaskliststatus" package.
	ProjectTaskListStatusesInverseTable = "project_task_list_status"
	// ProjectTaskListStatusesColumn is the table column denoting the project_task_list_statuses relation/edge.
	ProjectTaskListStatusesColumn = "task_list_completed_status_id"
)

// Columns holds all SQL columns for tasklistcompletedstatus fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldStatusCode,
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

// StatusCode defines the type for the "status_code" enum field.
type StatusCode string

// StatusCode values.
const (
	StatusCodeIncomplete         StatusCode = "INCOMPLETE"
	StatusCodeCompleted          StatusCode = "COMPLETED"
	StatusCodeCompletedToday     StatusCode = "COMPLETED_TODAY"
	StatusCodeCompletedYesterday StatusCode = "COMPLETED_YESTERDAY"
	StatusCodeCompleted1Week     StatusCode = "COMPLETED_1_WEEK"
	StatusCodeCompleted2Weeks    StatusCode = "COMPLETED_2_WEEKS"
	StatusCodeCompleted3Weeks    StatusCode = "COMPLETED_3_WEEKS"
	StatusCodeAll                StatusCode = "ALL"
)

func (sc StatusCode) String() string {
	return string(sc)
}

// StatusCodeValidator is a validator for the "status_code" field enum values. It is called by the builders before save.
func StatusCodeValidator(sc StatusCode) error {
	switch sc {
	case StatusCodeIncomplete, StatusCodeCompleted, StatusCodeCompletedToday, StatusCodeCompletedYesterday, StatusCodeCompleted1Week, StatusCodeCompleted2Weeks, StatusCodeCompleted3Weeks, StatusCodeAll:
		return nil
	default:
		return fmt.Errorf("tasklistcompletedstatus: invalid enum value for status_code field: %q", sc)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (sc StatusCode) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(sc.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (sc *StatusCode) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*sc = StatusCode(str)
	if err := StatusCodeValidator(*sc); err != nil {
		return fmt.Errorf("%s is not a valid StatusCode", str)
	}
	return nil
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/tasklistsortstatus"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TaskListSortStatus is the model entity for the TaskListSortStatus schema.
type TaskListSortStatus struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// StatusCode holds the value of the "status_code" field.
	StatusCode tasklistsortstatus.StatusCode `json:"status_code,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskListSortStatusQuery when eager-loading is set.
	Edges TaskListSortStatusEdges `json:"edges"`
}

// TaskListSortStatusEdges holds the relations/edges for other nodes in the graph.
type TaskListSortStatusEdges struct {
	// TeammateTaskListStatuses holds the value of the teammateTaskListStatuses edge.
	TeammateTaskListStatuses []*TeammateTaskListStatus `json:"teammateTaskListStatuses,omitempty"`
	// ProjectTaskListStatuses holds the value of the projectTaskListStatuses edge.
	ProjectTaskListStatuses []*ProjectTaskListStatus `json:"projectTaskListStatuses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TeammateTaskListStatusesOrErr returns the TeammateTaskListStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e TaskListSortStatusEdges) TeammateTaskListStatusesOrErr() ([]*TeammateTaskListStatus, error) {
	if e.loadedTypes[0] {
		return e.TeammateTaskListStatuses, nil
	}
	return nil, &NotLoadedError{edge: "teammateTaskListStatuses"}
}

// ProjectTaskListStatusesOrErr returns the ProjectTaskListStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e TaskListSortStatusEdges) ProjectTaskListStatusesOrErr() ([]*ProjectTaskListStatus, error) {
	if e.loadedTypes[1] {
		return e.ProjectTaskListStatuses, nil
	}
	return nil, &NotLoadedError{edge: "projectTaskListStatuses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TaskListSortStatus) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tasklistsortstatus.FieldName, tasklistsortstatus.FieldStatusCode:
			values[i] = new(sql.NullString)
		case tasklistsortstatus.FieldCreatedAt, tasklistsortstatus.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case tasklistsortstatus.FieldID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TaskListSortStatus", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TaskListSortStatus fields.
func (tlss *TaskListSortStatus) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tasklistsortstatus.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tlss.ID = *value
			}
		case tasklistsortstatus.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tlss.Name = value.String
			}
		case tasklistsortstatus.FieldStatusCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status_code", values[i])
			} else if value.Valid {
				tlss.StatusCode = tasklistsortstatus.StatusCode(value.String)
			}
		case tasklistsortstatus.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tlss.CreatedAt = value.Time
			}
		case tasklistsortstatus.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tlss.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTeammateTaskListStatuses queries the "teammateTaskListStatuses" edge of the TaskListSortStatus entity.
func (tlss *TaskListSortStatus) QueryTeammateTaskListStatuses() *TeammateTaskListStatusQuery {
	return (&TaskListSortStatusClient{config: tlss.config}).QueryTeammateTaskListStatuses(tlss)
}

// QueryProjectTaskListStatuses queries the "projectTaskListStatuses" edge of the TaskListSortStatus entity.
func (tlss *TaskListSortStatus) QueryProjectTaskListStatuses() *ProjectTaskListStatusQuery {
	return (&TaskListSortStatusClient{config: tlss.config}).QueryProjectTaskListStatuses(tlss)
}

// Update returns a builder for updating this TaskListSortStatus.
// Note that you need to call TaskListSortStatus.Unwrap() before calling this method if this TaskListSortStatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (tlss *TaskListSortStatus) Update() *TaskListSortStatusUpdateOne {
	return (&TaskListSortStatusClient{config: tlss.config}).UpdateOne(tlss)
}

// Unwrap unwraps the TaskListSortStatus entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tlss *TaskListSortStatus) Unwrap() *TaskListSortStatus {
	tx, ok := tlss.config.driver.(*txDriver)
	if !ok {
		panic("ent: TaskListSortStatus is not a transactional entity")
	}
	tlss.config.driver = tx.drv
	return tlss
}

// String implements the fmt.Stringer.
func (tlss *TaskListSortStatus) String() string {
	var builder strings.Builder
	builder.WriteString("TaskListSortStatus(")
	builder.WriteString(fmt.Sprintf("id=%v", tlss.ID))
	builder.WriteString(", name=")
	builder.WriteString(tlss.Name)
	builder.WriteString(", status_code=")
	builder.WriteString(fmt.Sprintf("%v", tlss.StatusCode))
	builder.WriteString(", created_at=")
	builder.WriteString(tlss.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(tlss.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TaskListSortStatusSlice is a parsable slice of TaskListSortStatus.
type TaskListSortStatusSlice []*TaskListSortStatus

func (tlss TaskListSortStatusSlice) config(cfg config) {
	for _i := range tlss {
		tlss[_i].config = cfg
	}
}

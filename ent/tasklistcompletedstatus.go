// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/tasklistcompletedstatus"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TaskListCompletedStatus is the model entity for the TaskListCompletedStatus schema.
type TaskListCompletedStatus struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// StatusCode holds the value of the "status_code" field.
	StatusCode tasklistcompletedstatus.StatusCode `json:"status_code,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TaskListCompletedStatus) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tasklistcompletedstatus.FieldName, tasklistcompletedstatus.FieldStatusCode:
			values[i] = new(sql.NullString)
		case tasklistcompletedstatus.FieldCreatedAt, tasklistcompletedstatus.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case tasklistcompletedstatus.FieldID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TaskListCompletedStatus", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TaskListCompletedStatus fields.
func (tlcs *TaskListCompletedStatus) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tasklistcompletedstatus.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tlcs.ID = *value
			}
		case tasklistcompletedstatus.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tlcs.Name = value.String
			}
		case tasklistcompletedstatus.FieldStatusCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status_code", values[i])
			} else if value.Valid {
				tlcs.StatusCode = tasklistcompletedstatus.StatusCode(value.String)
			}
		case tasklistcompletedstatus.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tlcs.CreatedAt = value.Time
			}
		case tasklistcompletedstatus.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tlcs.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TaskListCompletedStatus.
// Note that you need to call TaskListCompletedStatus.Unwrap() before calling this method if this TaskListCompletedStatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (tlcs *TaskListCompletedStatus) Update() *TaskListCompletedStatusUpdateOne {
	return (&TaskListCompletedStatusClient{config: tlcs.config}).UpdateOne(tlcs)
}

// Unwrap unwraps the TaskListCompletedStatus entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tlcs *TaskListCompletedStatus) Unwrap() *TaskListCompletedStatus {
	tx, ok := tlcs.config.driver.(*txDriver)
	if !ok {
		panic("ent: TaskListCompletedStatus is not a transactional entity")
	}
	tlcs.config.driver = tx.drv
	return tlcs
}

// String implements the fmt.Stringer.
func (tlcs *TaskListCompletedStatus) String() string {
	var builder strings.Builder
	builder.WriteString("TaskListCompletedStatus(")
	builder.WriteString(fmt.Sprintf("id=%v", tlcs.ID))
	builder.WriteString(", name=")
	builder.WriteString(tlcs.Name)
	builder.WriteString(", status_code=")
	builder.WriteString(fmt.Sprintf("%v", tlcs.StatusCode))
	builder.WriteString(", created_at=")
	builder.WriteString(tlcs.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(tlcs.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TaskListCompletedStatusSlice is a parsable slice of TaskListCompletedStatus.
type TaskListCompletedStatusSlice []*TaskListCompletedStatus

func (tlcs TaskListCompletedStatusSlice) config(cfg config) {
	for _i := range tlcs {
		tlcs[_i].config = cfg
	}
}

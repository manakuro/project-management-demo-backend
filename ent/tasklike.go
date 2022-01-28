// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/tasklike"
	"project-management-demo-backend/ent/teammate"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TaskLike is the model entity for the TaskLike schema.
type TaskLike struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// TaskID holds the value of the "task_id" field.
	TaskID ulid.ID `json:"task_id,omitempty"`
	// TeammateID holds the value of the "teammate_id" field.
	TeammateID ulid.ID `json:"teammate_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskLikeQuery when eager-loading is set.
	Edges TaskLikeEdges `json:"edges"`
}

// TaskLikeEdges holds the relations/edges for other nodes in the graph.
type TaskLikeEdges struct {
	// Task holds the value of the task edge.
	Task *Task `json:"task,omitempty"`
	// Teammate holds the value of the teammate edge.
	Teammate *Teammate `json:"teammate,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskLikeEdges) TaskOrErr() (*Task, error) {
	if e.loadedTypes[0] {
		if e.Task == nil {
			// The edge task was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: task.Label}
		}
		return e.Task, nil
	}
	return nil, &NotLoadedError{edge: "task"}
}

// TeammateOrErr returns the Teammate value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskLikeEdges) TeammateOrErr() (*Teammate, error) {
	if e.loadedTypes[1] {
		if e.Teammate == nil {
			// The edge teammate was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: teammate.Label}
		}
		return e.Teammate, nil
	}
	return nil, &NotLoadedError{edge: "teammate"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TaskLike) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tasklike.FieldCreatedAt, tasklike.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case tasklike.FieldID, tasklike.FieldTaskID, tasklike.FieldTeammateID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TaskLike", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TaskLike fields.
func (tl *TaskLike) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tasklike.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tl.ID = *value
			}
		case tasklike.FieldTaskID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field task_id", values[i])
			} else if value != nil {
				tl.TaskID = *value
			}
		case tasklike.FieldTeammateID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field teammate_id", values[i])
			} else if value != nil {
				tl.TeammateID = *value
			}
		case tasklike.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tl.CreatedAt = value.Time
			}
		case tasklike.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tl.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTask queries the "task" edge of the TaskLike entity.
func (tl *TaskLike) QueryTask() *TaskQuery {
	return (&TaskLikeClient{config: tl.config}).QueryTask(tl)
}

// QueryTeammate queries the "teammate" edge of the TaskLike entity.
func (tl *TaskLike) QueryTeammate() *TeammateQuery {
	return (&TaskLikeClient{config: tl.config}).QueryTeammate(tl)
}

// Update returns a builder for updating this TaskLike.
// Note that you need to call TaskLike.Unwrap() before calling this method if this TaskLike
// was returned from a transaction, and the transaction was committed or rolled back.
func (tl *TaskLike) Update() *TaskLikeUpdateOne {
	return (&TaskLikeClient{config: tl.config}).UpdateOne(tl)
}

// Unwrap unwraps the TaskLike entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tl *TaskLike) Unwrap() *TaskLike {
	tx, ok := tl.config.driver.(*txDriver)
	if !ok {
		panic("ent: TaskLike is not a transactional entity")
	}
	tl.config.driver = tx.drv
	return tl
}

// String implements the fmt.Stringer.
func (tl *TaskLike) String() string {
	var builder strings.Builder
	builder.WriteString("TaskLike(")
	builder.WriteString(fmt.Sprintf("id=%v", tl.ID))
	builder.WriteString(", task_id=")
	builder.WriteString(fmt.Sprintf("%v", tl.TaskID))
	builder.WriteString(", teammate_id=")
	builder.WriteString(fmt.Sprintf("%v", tl.TeammateID))
	builder.WriteString(", created_at=")
	builder.WriteString(tl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(tl.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TaskLikes is a parsable slice of TaskLike.
type TaskLikes []*TaskLike

func (tl TaskLikes) config(cfg config) {
	for _i := range tl {
		tl[_i].config = cfg
	}
}

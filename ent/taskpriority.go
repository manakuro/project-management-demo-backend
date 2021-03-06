// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/color"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/taskpriority"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TaskPriority is the model entity for the TaskPriority schema.
type TaskPriority struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// ColorID holds the value of the "color_id" field.
	ColorID ulid.ID `json:"color_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// PriorityType holds the value of the "priority_type" field.
	PriorityType taskpriority.PriorityType `json:"priority_type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskPriorityQuery when eager-loading is set.
	Edges TaskPriorityEdges `json:"edges"`
}

// TaskPriorityEdges holds the relations/edges for other nodes in the graph.
type TaskPriorityEdges struct {
	// Color holds the value of the color edge.
	Color *Color `json:"color,omitempty"`
	// Tasks holds the value of the tasks edge.
	Tasks []*Task `json:"tasks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ColorOrErr returns the Color value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskPriorityEdges) ColorOrErr() (*Color, error) {
	if e.loadedTypes[0] {
		if e.Color == nil {
			// The edge color was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: color.Label}
		}
		return e.Color, nil
	}
	return nil, &NotLoadedError{edge: "color"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e TaskPriorityEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[1] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TaskPriority) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case taskpriority.FieldName, taskpriority.FieldPriorityType:
			values[i] = new(sql.NullString)
		case taskpriority.FieldCreatedAt, taskpriority.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case taskpriority.FieldID, taskpriority.FieldColorID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TaskPriority", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TaskPriority fields.
func (tp *TaskPriority) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case taskpriority.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tp.ID = *value
			}
		case taskpriority.FieldColorID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field color_id", values[i])
			} else if value != nil {
				tp.ColorID = *value
			}
		case taskpriority.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tp.Name = value.String
			}
		case taskpriority.FieldPriorityType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field priority_type", values[i])
			} else if value.Valid {
				tp.PriorityType = taskpriority.PriorityType(value.String)
			}
		case taskpriority.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tp.CreatedAt = value.Time
			}
		case taskpriority.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tp.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryColor queries the "color" edge of the TaskPriority entity.
func (tp *TaskPriority) QueryColor() *ColorQuery {
	return (&TaskPriorityClient{config: tp.config}).QueryColor(tp)
}

// QueryTasks queries the "tasks" edge of the TaskPriority entity.
func (tp *TaskPriority) QueryTasks() *TaskQuery {
	return (&TaskPriorityClient{config: tp.config}).QueryTasks(tp)
}

// Update returns a builder for updating this TaskPriority.
// Note that you need to call TaskPriority.Unwrap() before calling this method if this TaskPriority
// was returned from a transaction, and the transaction was committed or rolled back.
func (tp *TaskPriority) Update() *TaskPriorityUpdateOne {
	return (&TaskPriorityClient{config: tp.config}).UpdateOne(tp)
}

// Unwrap unwraps the TaskPriority entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tp *TaskPriority) Unwrap() *TaskPriority {
	tx, ok := tp.config.driver.(*txDriver)
	if !ok {
		panic("ent: TaskPriority is not a transactional entity")
	}
	tp.config.driver = tx.drv
	return tp
}

// String implements the fmt.Stringer.
func (tp *TaskPriority) String() string {
	var builder strings.Builder
	builder.WriteString("TaskPriority(")
	builder.WriteString(fmt.Sprintf("id=%v", tp.ID))
	builder.WriteString(", color_id=")
	builder.WriteString(fmt.Sprintf("%v", tp.ColorID))
	builder.WriteString(", name=")
	builder.WriteString(tp.Name)
	builder.WriteString(", priority_type=")
	builder.WriteString(fmt.Sprintf("%v", tp.PriorityType))
	builder.WriteString(", created_at=")
	builder.WriteString(tp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(tp.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TaskPriorities is a parsable slice of TaskPriority.
type TaskPriorities []*TaskPriority

func (tp TaskPriorities) config(cfg config) {
	for _i := range tp {
		tp[_i].config = cfg
	}
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/deletedtask"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/workspace"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// DeletedTask is the model entity for the DeletedTask schema.
type DeletedTask struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// TaskID holds the value of the "task_id" field.
	TaskID ulid.ID `json:"task_id,omitempty"`
	// WorkspaceID holds the value of the "workspace_id" field.
	WorkspaceID ulid.ID `json:"workspace_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeletedTaskQuery when eager-loading is set.
	Edges DeletedTaskEdges `json:"edges"`
}

// DeletedTaskEdges holds the relations/edges for other nodes in the graph.
type DeletedTaskEdges struct {
	// Task holds the value of the task edge.
	Task *Task `json:"task,omitempty"`
	// Workspace holds the value of the workspace edge.
	Workspace *Workspace `json:"workspace,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeletedTaskEdges) TaskOrErr() (*Task, error) {
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

// WorkspaceOrErr returns the Workspace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeletedTaskEdges) WorkspaceOrErr() (*Workspace, error) {
	if e.loadedTypes[1] {
		if e.Workspace == nil {
			// The edge workspace was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workspace.Label}
		}
		return e.Workspace, nil
	}
	return nil, &NotLoadedError{edge: "workspace"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DeletedTask) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case deletedtask.FieldCreatedAt, deletedtask.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case deletedtask.FieldID, deletedtask.FieldTaskID, deletedtask.FieldWorkspaceID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DeletedTask", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeletedTask fields.
func (dt *DeletedTask) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deletedtask.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				dt.ID = *value
			}
		case deletedtask.FieldTaskID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field task_id", values[i])
			} else if value != nil {
				dt.TaskID = *value
			}
		case deletedtask.FieldWorkspaceID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field workspace_id", values[i])
			} else if value != nil {
				dt.WorkspaceID = *value
			}
		case deletedtask.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dt.CreatedAt = value.Time
			}
		case deletedtask.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dt.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTask queries the "task" edge of the DeletedTask entity.
func (dt *DeletedTask) QueryTask() *TaskQuery {
	return (&DeletedTaskClient{config: dt.config}).QueryTask(dt)
}

// QueryWorkspace queries the "workspace" edge of the DeletedTask entity.
func (dt *DeletedTask) QueryWorkspace() *WorkspaceQuery {
	return (&DeletedTaskClient{config: dt.config}).QueryWorkspace(dt)
}

// Update returns a builder for updating this DeletedTask.
// Note that you need to call DeletedTask.Unwrap() before calling this method if this DeletedTask
// was returned from a transaction, and the transaction was committed or rolled back.
func (dt *DeletedTask) Update() *DeletedTaskUpdateOne {
	return (&DeletedTaskClient{config: dt.config}).UpdateOne(dt)
}

// Unwrap unwraps the DeletedTask entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dt *DeletedTask) Unwrap() *DeletedTask {
	tx, ok := dt.config.driver.(*txDriver)
	if !ok {
		panic("ent: DeletedTask is not a transactional entity")
	}
	dt.config.driver = tx.drv
	return dt
}

// String implements the fmt.Stringer.
func (dt *DeletedTask) String() string {
	var builder strings.Builder
	builder.WriteString("DeletedTask(")
	builder.WriteString(fmt.Sprintf("id=%v", dt.ID))
	builder.WriteString(", task_id=")
	builder.WriteString(fmt.Sprintf("%v", dt.TaskID))
	builder.WriteString(", workspace_id=")
	builder.WriteString(fmt.Sprintf("%v", dt.WorkspaceID))
	builder.WriteString(", created_at=")
	builder.WriteString(dt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(dt.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// DeletedTasks is a parsable slice of DeletedTask.
type DeletedTasks []*DeletedTask

func (dt DeletedTasks) config(cfg config) {
	for _i := range dt {
		dt[_i].config = cfg
	}
}

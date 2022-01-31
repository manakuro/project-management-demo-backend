// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/filetype"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/taskfeed"
	"project-management-demo-backend/ent/taskfile"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TaskFile is the model entity for the TaskFile schema.
type TaskFile struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID ulid.ID `json:"project_id,omitempty"`
	// TaskID holds the value of the "task_id" field.
	TaskID ulid.ID `json:"task_id,omitempty"`
	// TaskFeedID holds the value of the "task_feed_id" field.
	TaskFeedID ulid.ID `json:"task_feed_id,omitempty"`
	// FileTypeID holds the value of the "file_type_id" field.
	FileTypeID ulid.ID `json:"file_type_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Src holds the value of the "src" field.
	Src string `json:"src,omitempty"`
	// Attached holds the value of the "attached" field.
	Attached bool `json:"attached,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskFileQuery when eager-loading is set.
	Edges TaskFileEdges `json:"edges"`
}

// TaskFileEdges holds the relations/edges for other nodes in the graph.
type TaskFileEdges struct {
	// Project holds the value of the project edge.
	Project *Project `json:"project,omitempty"`
	// Task holds the value of the task edge.
	Task *Task `json:"task,omitempty"`
	// TaskFeed holds the value of the task_feed edge.
	TaskFeed *TaskFeed `json:"task_feed,omitempty"`
	// FileType holds the value of the file_type edge.
	FileType *FileType `json:"file_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskFileEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// The edge project was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskFileEdges) TaskOrErr() (*Task, error) {
	if e.loadedTypes[1] {
		if e.Task == nil {
			// The edge task was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: task.Label}
		}
		return e.Task, nil
	}
	return nil, &NotLoadedError{edge: "task"}
}

// TaskFeedOrErr returns the TaskFeed value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskFileEdges) TaskFeedOrErr() (*TaskFeed, error) {
	if e.loadedTypes[2] {
		if e.TaskFeed == nil {
			// The edge task_feed was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: taskfeed.Label}
		}
		return e.TaskFeed, nil
	}
	return nil, &NotLoadedError{edge: "task_feed"}
}

// FileTypeOrErr returns the FileType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskFileEdges) FileTypeOrErr() (*FileType, error) {
	if e.loadedTypes[3] {
		if e.FileType == nil {
			// The edge file_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: filetype.Label}
		}
		return e.FileType, nil
	}
	return nil, &NotLoadedError{edge: "file_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TaskFile) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case taskfile.FieldAttached:
			values[i] = new(sql.NullBool)
		case taskfile.FieldName, taskfile.FieldSrc:
			values[i] = new(sql.NullString)
		case taskfile.FieldCreatedAt, taskfile.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case taskfile.FieldID, taskfile.FieldProjectID, taskfile.FieldTaskID, taskfile.FieldTaskFeedID, taskfile.FieldFileTypeID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TaskFile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TaskFile fields.
func (tf *TaskFile) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case taskfile.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tf.ID = *value
			}
		case taskfile.FieldProjectID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				tf.ProjectID = *value
			}
		case taskfile.FieldTaskID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field task_id", values[i])
			} else if value != nil {
				tf.TaskID = *value
			}
		case taskfile.FieldTaskFeedID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field task_feed_id", values[i])
			} else if value != nil {
				tf.TaskFeedID = *value
			}
		case taskfile.FieldFileTypeID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field file_type_id", values[i])
			} else if value != nil {
				tf.FileTypeID = *value
			}
		case taskfile.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tf.Name = value.String
			}
		case taskfile.FieldSrc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field src", values[i])
			} else if value.Valid {
				tf.Src = value.String
			}
		case taskfile.FieldAttached:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field attached", values[i])
			} else if value.Valid {
				tf.Attached = value.Bool
			}
		case taskfile.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tf.CreatedAt = value.Time
			}
		case taskfile.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tf.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryProject queries the "project" edge of the TaskFile entity.
func (tf *TaskFile) QueryProject() *ProjectQuery {
	return (&TaskFileClient{config: tf.config}).QueryProject(tf)
}

// QueryTask queries the "task" edge of the TaskFile entity.
func (tf *TaskFile) QueryTask() *TaskQuery {
	return (&TaskFileClient{config: tf.config}).QueryTask(tf)
}

// QueryTaskFeed queries the "task_feed" edge of the TaskFile entity.
func (tf *TaskFile) QueryTaskFeed() *TaskFeedQuery {
	return (&TaskFileClient{config: tf.config}).QueryTaskFeed(tf)
}

// QueryFileType queries the "file_type" edge of the TaskFile entity.
func (tf *TaskFile) QueryFileType() *FileTypeQuery {
	return (&TaskFileClient{config: tf.config}).QueryFileType(tf)
}

// Update returns a builder for updating this TaskFile.
// Note that you need to call TaskFile.Unwrap() before calling this method if this TaskFile
// was returned from a transaction, and the transaction was committed or rolled back.
func (tf *TaskFile) Update() *TaskFileUpdateOne {
	return (&TaskFileClient{config: tf.config}).UpdateOne(tf)
}

// Unwrap unwraps the TaskFile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tf *TaskFile) Unwrap() *TaskFile {
	tx, ok := tf.config.driver.(*txDriver)
	if !ok {
		panic("ent: TaskFile is not a transactional entity")
	}
	tf.config.driver = tx.drv
	return tf
}

// String implements the fmt.Stringer.
func (tf *TaskFile) String() string {
	var builder strings.Builder
	builder.WriteString("TaskFile(")
	builder.WriteString(fmt.Sprintf("id=%v", tf.ID))
	builder.WriteString(", project_id=")
	builder.WriteString(fmt.Sprintf("%v", tf.ProjectID))
	builder.WriteString(", task_id=")
	builder.WriteString(fmt.Sprintf("%v", tf.TaskID))
	builder.WriteString(", task_feed_id=")
	builder.WriteString(fmt.Sprintf("%v", tf.TaskFeedID))
	builder.WriteString(", file_type_id=")
	builder.WriteString(fmt.Sprintf("%v", tf.FileTypeID))
	builder.WriteString(", name=")
	builder.WriteString(tf.Name)
	builder.WriteString(", src=")
	builder.WriteString(tf.Src)
	builder.WriteString(", attached=")
	builder.WriteString(fmt.Sprintf("%v", tf.Attached))
	builder.WriteString(", created_at=")
	builder.WriteString(tf.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(tf.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TaskFiles is a parsable slice of TaskFile.
type TaskFiles []*TaskFile

func (tf TaskFiles) config(cfg config) {
	for _i := range tf {
		tf[_i].config = cfg
	}
}

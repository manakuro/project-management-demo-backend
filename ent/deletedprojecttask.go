// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/deletedprojecttask"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// DeletedProjectTask is the model entity for the DeletedProjectTask schema.
type DeletedProjectTask struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID ulid.ID `json:"project_id,omitempty"`
	// TaskID holds the value of the "task_id" field.
	TaskID ulid.ID `json:"task_id,omitempty"`
	// ProjectTaskSectionID holds the value of the "project_task_section_id" field.
	ProjectTaskSectionID ulid.ID `json:"project_task_section_id,omitempty"`
	// ProjectTaskID holds the value of the "project_task_id" field.
	ProjectTaskID ulid.ID `json:"project_task_id,omitempty"`
	// ProjectTaskCreatedAt holds the value of the "project_task_created_at" field.
	ProjectTaskCreatedAt time.Time `json:"project_task_created_at,omitempty"`
	// ProjectTaskUpdatedAt holds the value of the "project_task_updated_at" field.
	ProjectTaskUpdatedAt time.Time `json:"project_task_updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeletedProjectTaskQuery when eager-loading is set.
	Edges DeletedProjectTaskEdges `json:"edges"`
}

// DeletedProjectTaskEdges holds the relations/edges for other nodes in the graph.
type DeletedProjectTaskEdges struct {
	// Project holds the value of the project edge.
	Project *Project `json:"project,omitempty"`
	// Task holds the value of the task edge.
	Task *Task `json:"task,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeletedProjectTaskEdges) ProjectOrErr() (*Project, error) {
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
func (e DeletedProjectTaskEdges) TaskOrErr() (*Task, error) {
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

// scanValues returns the types for scanning values from sql.Rows.
func (*DeletedProjectTask) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case deletedprojecttask.FieldProjectTaskCreatedAt, deletedprojecttask.FieldProjectTaskUpdatedAt, deletedprojecttask.FieldCreatedAt, deletedprojecttask.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case deletedprojecttask.FieldID, deletedprojecttask.FieldProjectID, deletedprojecttask.FieldTaskID, deletedprojecttask.FieldProjectTaskSectionID, deletedprojecttask.FieldProjectTaskID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DeletedProjectTask", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeletedProjectTask fields.
func (dpt *DeletedProjectTask) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deletedprojecttask.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				dpt.ID = *value
			}
		case deletedprojecttask.FieldProjectID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				dpt.ProjectID = *value
			}
		case deletedprojecttask.FieldTaskID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field task_id", values[i])
			} else if value != nil {
				dpt.TaskID = *value
			}
		case deletedprojecttask.FieldProjectTaskSectionID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_task_section_id", values[i])
			} else if value != nil {
				dpt.ProjectTaskSectionID = *value
			}
		case deletedprojecttask.FieldProjectTaskID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_task_id", values[i])
			} else if value != nil {
				dpt.ProjectTaskID = *value
			}
		case deletedprojecttask.FieldProjectTaskCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field project_task_created_at", values[i])
			} else if value.Valid {
				dpt.ProjectTaskCreatedAt = value.Time
			}
		case deletedprojecttask.FieldProjectTaskUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field project_task_updated_at", values[i])
			} else if value.Valid {
				dpt.ProjectTaskUpdatedAt = value.Time
			}
		case deletedprojecttask.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dpt.CreatedAt = value.Time
			}
		case deletedprojecttask.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dpt.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryProject queries the "project" edge of the DeletedProjectTask entity.
func (dpt *DeletedProjectTask) QueryProject() *ProjectQuery {
	return (&DeletedProjectTaskClient{config: dpt.config}).QueryProject(dpt)
}

// QueryTask queries the "task" edge of the DeletedProjectTask entity.
func (dpt *DeletedProjectTask) QueryTask() *TaskQuery {
	return (&DeletedProjectTaskClient{config: dpt.config}).QueryTask(dpt)
}

// Update returns a builder for updating this DeletedProjectTask.
// Note that you need to call DeletedProjectTask.Unwrap() before calling this method if this DeletedProjectTask
// was returned from a transaction, and the transaction was committed or rolled back.
func (dpt *DeletedProjectTask) Update() *DeletedProjectTaskUpdateOne {
	return (&DeletedProjectTaskClient{config: dpt.config}).UpdateOne(dpt)
}

// Unwrap unwraps the DeletedProjectTask entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dpt *DeletedProjectTask) Unwrap() *DeletedProjectTask {
	tx, ok := dpt.config.driver.(*txDriver)
	if !ok {
		panic("ent: DeletedProjectTask is not a transactional entity")
	}
	dpt.config.driver = tx.drv
	return dpt
}

// String implements the fmt.Stringer.
func (dpt *DeletedProjectTask) String() string {
	var builder strings.Builder
	builder.WriteString("DeletedProjectTask(")
	builder.WriteString(fmt.Sprintf("id=%v", dpt.ID))
	builder.WriteString(", project_id=")
	builder.WriteString(fmt.Sprintf("%v", dpt.ProjectID))
	builder.WriteString(", task_id=")
	builder.WriteString(fmt.Sprintf("%v", dpt.TaskID))
	builder.WriteString(", project_task_section_id=")
	builder.WriteString(fmt.Sprintf("%v", dpt.ProjectTaskSectionID))
	builder.WriteString(", project_task_id=")
	builder.WriteString(fmt.Sprintf("%v", dpt.ProjectTaskID))
	builder.WriteString(", project_task_created_at=")
	builder.WriteString(dpt.ProjectTaskCreatedAt.Format(time.ANSIC))
	builder.WriteString(", project_task_updated_at=")
	builder.WriteString(dpt.ProjectTaskUpdatedAt.Format(time.ANSIC))
	builder.WriteString(", created_at=")
	builder.WriteString(dpt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(dpt.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// DeletedProjectTasks is a parsable slice of DeletedProjectTask.
type DeletedProjectTasks []*DeletedProjectTask

func (dpt DeletedProjectTasks) config(cfg config) {
	for _i := range dpt {
		dpt[_i].config = cfg
	}
}

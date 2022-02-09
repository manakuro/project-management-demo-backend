// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/project"
	"project-management-demo-backend/ent/projecttask"
	"project-management-demo-backend/ent/projecttasksection"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ProjectTask is the model entity for the ProjectTask schema.
type ProjectTask struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID ulid.ID `json:"project_id,omitempty"`
	// TaskID holds the value of the "task_id" field.
	TaskID ulid.ID `json:"task_id,omitempty"`
	// ProjectTaskSectionID holds the value of the "project_task_section_id" field.
	ProjectTaskSectionID ulid.ID `json:"project_task_section_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectTaskQuery when eager-loading is set.
	Edges ProjectTaskEdges `json:"edges"`
}

// ProjectTaskEdges holds the relations/edges for other nodes in the graph.
type ProjectTaskEdges struct {
	// Project holds the value of the project edge.
	Project *Project `json:"project,omitempty"`
	// Task holds the value of the task edge.
	Task *Task `json:"task,omitempty"`
	// ProjectTaskSection holds the value of the projectTaskSection edge.
	ProjectTaskSection *ProjectTaskSection `json:"projectTaskSection,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectTaskEdges) ProjectOrErr() (*Project, error) {
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
func (e ProjectTaskEdges) TaskOrErr() (*Task, error) {
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

// ProjectTaskSectionOrErr returns the ProjectTaskSection value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectTaskEdges) ProjectTaskSectionOrErr() (*ProjectTaskSection, error) {
	if e.loadedTypes[2] {
		if e.ProjectTaskSection == nil {
			// The edge projectTaskSection was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: projecttasksection.Label}
		}
		return e.ProjectTaskSection, nil
	}
	return nil, &NotLoadedError{edge: "projectTaskSection"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProjectTask) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case projecttask.FieldCreatedAt, projecttask.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case projecttask.FieldID, projecttask.FieldProjectID, projecttask.FieldTaskID, projecttask.FieldProjectTaskSectionID:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProjectTask", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProjectTask fields.
func (pt *ProjectTask) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case projecttask.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pt.ID = *value
			}
		case projecttask.FieldProjectID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				pt.ProjectID = *value
			}
		case projecttask.FieldTaskID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field task_id", values[i])
			} else if value != nil {
				pt.TaskID = *value
			}
		case projecttask.FieldProjectTaskSectionID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_task_section_id", values[i])
			} else if value != nil {
				pt.ProjectTaskSectionID = *value
			}
		case projecttask.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pt.CreatedAt = value.Time
			}
		case projecttask.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pt.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryProject queries the "project" edge of the ProjectTask entity.
func (pt *ProjectTask) QueryProject() *ProjectQuery {
	return (&ProjectTaskClient{config: pt.config}).QueryProject(pt)
}

// QueryTask queries the "task" edge of the ProjectTask entity.
func (pt *ProjectTask) QueryTask() *TaskQuery {
	return (&ProjectTaskClient{config: pt.config}).QueryTask(pt)
}

// QueryProjectTaskSection queries the "projectTaskSection" edge of the ProjectTask entity.
func (pt *ProjectTask) QueryProjectTaskSection() *ProjectTaskSectionQuery {
	return (&ProjectTaskClient{config: pt.config}).QueryProjectTaskSection(pt)
}

// Update returns a builder for updating this ProjectTask.
// Note that you need to call ProjectTask.Unwrap() before calling this method if this ProjectTask
// was returned from a transaction, and the transaction was committed or rolled back.
func (pt *ProjectTask) Update() *ProjectTaskUpdateOne {
	return (&ProjectTaskClient{config: pt.config}).UpdateOne(pt)
}

// Unwrap unwraps the ProjectTask entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pt *ProjectTask) Unwrap() *ProjectTask {
	tx, ok := pt.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProjectTask is not a transactional entity")
	}
	pt.config.driver = tx.drv
	return pt
}

// String implements the fmt.Stringer.
func (pt *ProjectTask) String() string {
	var builder strings.Builder
	builder.WriteString("ProjectTask(")
	builder.WriteString(fmt.Sprintf("id=%v", pt.ID))
	builder.WriteString(", project_id=")
	builder.WriteString(fmt.Sprintf("%v", pt.ProjectID))
	builder.WriteString(", task_id=")
	builder.WriteString(fmt.Sprintf("%v", pt.TaskID))
	builder.WriteString(", project_task_section_id=")
	builder.WriteString(fmt.Sprintf("%v", pt.ProjectTaskSectionID))
	builder.WriteString(", created_at=")
	builder.WriteString(pt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(pt.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ProjectTasks is a parsable slice of ProjectTask.
type ProjectTasks []*ProjectTask

func (pt ProjectTasks) config(cfg config) {
	for _i := range pt {
		pt[_i].config = cfg
	}
}

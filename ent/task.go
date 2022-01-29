// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/task"
	"project-management-demo-backend/ent/taskpriority"
	"project-management-demo-backend/ent/teammate"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// TaskParentID holds the value of the "task_parent_id" field.
	TaskParentID ulid.ID `json:"task_parent_id,omitempty"`
	// TaskPriorityID holds the value of the "task_priority_id" field.
	TaskPriorityID ulid.ID `json:"task_priority_id,omitempty"`
	// AssigneeID holds the value of the "assignee_id" field.
	AssigneeID ulid.ID `json:"assignee_id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy ulid.ID `json:"created_by,omitempty"`
	// Completed holds the value of the "completed" field.
	Completed bool `json:"completed,omitempty"`
	// CompletedAt holds the value of the "completed_at" field.
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// IsNew holds the value of the "is_new" field.
	IsNew bool `json:"is_new,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DueDate holds the value of the "due_date" field.
	DueDate *time.Time `json:"due_date,omitempty"`
	// DueTime holds the value of the "due_time" field.
	DueTime *time.Time `json:"due_time,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskQuery when eager-loading is set.
	Edges TaskEdges `json:"edges"`
}

// TaskEdges holds the relations/edges for other nodes in the graph.
type TaskEdges struct {
	// Teammate holds the value of the teammate edge.
	Teammate *Teammate `json:"teammate,omitempty"`
	// TaskPriority holds the value of the task_priority edge.
	TaskPriority *TaskPriority `json:"task_priority,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Task `json:"parent,omitempty"`
	// SubTasks holds the value of the sub_tasks edge.
	SubTasks []*Task `json:"sub_tasks,omitempty"`
	// TeammateTasks holds the value of the teammate_tasks edge.
	TeammateTasks []*TeammateTask `json:"teammate_tasks,omitempty"`
	// ProjectTasks holds the value of the project_tasks edge.
	ProjectTasks []*ProjectTask `json:"project_tasks,omitempty"`
	// TaskLikes holds the value of the task_likes edge.
	TaskLikes []*TaskLike `json:"task_likes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// TeammateOrErr returns the Teammate value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) TeammateOrErr() (*Teammate, error) {
	if e.loadedTypes[0] {
		if e.Teammate == nil {
			// The edge teammate was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: teammate.Label}
		}
		return e.Teammate, nil
	}
	return nil, &NotLoadedError{edge: "teammate"}
}

// TaskPriorityOrErr returns the TaskPriority value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) TaskPriorityOrErr() (*TaskPriority, error) {
	if e.loadedTypes[1] {
		if e.TaskPriority == nil {
			// The edge task_priority was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: taskpriority.Label}
		}
		return e.TaskPriority, nil
	}
	return nil, &NotLoadedError{edge: "task_priority"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) ParentOrErr() (*Task, error) {
	if e.loadedTypes[2] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: task.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// SubTasksOrErr returns the SubTasks value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) SubTasksOrErr() ([]*Task, error) {
	if e.loadedTypes[3] {
		return e.SubTasks, nil
	}
	return nil, &NotLoadedError{edge: "sub_tasks"}
}

// TeammateTasksOrErr returns the TeammateTasks value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) TeammateTasksOrErr() ([]*TeammateTask, error) {
	if e.loadedTypes[4] {
		return e.TeammateTasks, nil
	}
	return nil, &NotLoadedError{edge: "teammate_tasks"}
}

// ProjectTasksOrErr returns the ProjectTasks value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) ProjectTasksOrErr() ([]*ProjectTask, error) {
	if e.loadedTypes[5] {
		return e.ProjectTasks, nil
	}
	return nil, &NotLoadedError{edge: "project_tasks"}
}

// TaskLikesOrErr returns the TaskLikes value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) TaskLikesOrErr() ([]*TaskLike, error) {
	if e.loadedTypes[6] {
		return e.TaskLikes, nil
	}
	return nil, &NotLoadedError{edge: "task_likes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case task.FieldCompleted, task.FieldIsNew:
			values[i] = new(sql.NullBool)
		case task.FieldName:
			values[i] = new(sql.NullString)
		case task.FieldCompletedAt, task.FieldDueDate, task.FieldDueTime, task.FieldCreatedAt, task.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case task.FieldID, task.FieldTaskParentID, task.FieldTaskPriorityID, task.FieldAssigneeID, task.FieldCreatedBy:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Task", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case task.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case task.FieldTaskParentID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field task_parent_id", values[i])
			} else if value != nil {
				t.TaskParentID = *value
			}
		case task.FieldTaskPriorityID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field task_priority_id", values[i])
			} else if value != nil {
				t.TaskPriorityID = *value
			}
		case task.FieldAssigneeID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field assignee_id", values[i])
			} else if value != nil {
				t.AssigneeID = *value
			}
		case task.FieldCreatedBy:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				t.CreatedBy = *value
			}
		case task.FieldCompleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field completed", values[i])
			} else if value.Valid {
				t.Completed = value.Bool
			}
		case task.FieldCompletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field completed_at", values[i])
			} else if value.Valid {
				t.CompletedAt = new(time.Time)
				*t.CompletedAt = value.Time
			}
		case task.FieldIsNew:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_new", values[i])
			} else if value.Valid {
				t.IsNew = value.Bool
			}
		case task.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case task.FieldDueDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field due_date", values[i])
			} else if value.Valid {
				t.DueDate = new(time.Time)
				*t.DueDate = value.Time
			}
		case task.FieldDueTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field due_time", values[i])
			} else if value.Valid {
				t.DueTime = new(time.Time)
				*t.DueTime = value.Time
			}
		case task.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case task.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTeammate queries the "teammate" edge of the Task entity.
func (t *Task) QueryTeammate() *TeammateQuery {
	return (&TaskClient{config: t.config}).QueryTeammate(t)
}

// QueryTaskPriority queries the "task_priority" edge of the Task entity.
func (t *Task) QueryTaskPriority() *TaskPriorityQuery {
	return (&TaskClient{config: t.config}).QueryTaskPriority(t)
}

// QueryParent queries the "parent" edge of the Task entity.
func (t *Task) QueryParent() *TaskQuery {
	return (&TaskClient{config: t.config}).QueryParent(t)
}

// QuerySubTasks queries the "sub_tasks" edge of the Task entity.
func (t *Task) QuerySubTasks() *TaskQuery {
	return (&TaskClient{config: t.config}).QuerySubTasks(t)
}

// QueryTeammateTasks queries the "teammate_tasks" edge of the Task entity.
func (t *Task) QueryTeammateTasks() *TeammateTaskQuery {
	return (&TaskClient{config: t.config}).QueryTeammateTasks(t)
}

// QueryProjectTasks queries the "project_tasks" edge of the Task entity.
func (t *Task) QueryProjectTasks() *ProjectTaskQuery {
	return (&TaskClient{config: t.config}).QueryProjectTasks(t)
}

// QueryTaskLikes queries the "task_likes" edge of the Task entity.
func (t *Task) QueryTaskLikes() *TaskLikeQuery {
	return (&TaskClient{config: t.config}).QueryTaskLikes(t)
}

// Update returns a builder for updating this Task.
// Note that you need to call Task.Unwrap() before calling this method if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return (&TaskClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Task entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", task_parent_id=")
	builder.WriteString(fmt.Sprintf("%v", t.TaskParentID))
	builder.WriteString(", task_priority_id=")
	builder.WriteString(fmt.Sprintf("%v", t.TaskPriorityID))
	builder.WriteString(", assignee_id=")
	builder.WriteString(fmt.Sprintf("%v", t.AssigneeID))
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", t.CreatedBy))
	builder.WriteString(", completed=")
	builder.WriteString(fmt.Sprintf("%v", t.Completed))
	if v := t.CompletedAt; v != nil {
		builder.WriteString(", completed_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", is_new=")
	builder.WriteString(fmt.Sprintf("%v", t.IsNew))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	if v := t.DueDate; v != nil {
		builder.WriteString(", due_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := t.DueTime; v != nil {
		builder.WriteString(", due_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task

func (t Tasks) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/ent/teammate"
	"project-management-demo-backend/ent/workspace"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Workspace is the model entity for the Workspace schema.
type Workspace struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy ulid.ID `json:"created_by,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description map[string]interface{} `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkspaceQuery when eager-loading is set.
	Edges WorkspaceEdges `json:"edges"`
}

// WorkspaceEdges holds the relations/edges for other nodes in the graph.
type WorkspaceEdges struct {
	// Teammate holds the value of the teammate edge.
	Teammate *Teammate `json:"teammate,omitempty"`
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// WorkspaceTeammates holds the value of the workspaceTeammates edge.
	WorkspaceTeammates []*WorkspaceTeammate `json:"workspaceTeammates,omitempty"`
	// FavoriteWorkspaces holds the value of the favoriteWorkspaces edge.
	FavoriteWorkspaces []*FavoriteWorkspace `json:"favoriteWorkspaces,omitempty"`
	// TeammateTaskTabStatuses holds the value of the teammateTaskTabStatuses edge.
	TeammateTaskTabStatuses []*TeammateTaskTabStatus `json:"teammateTaskTabStatuses,omitempty"`
	// TeammateTaskListStatuses holds the value of the teammateTaskListStatuses edge.
	TeammateTaskListStatuses []*TeammateTaskListStatus `json:"teammateTaskListStatuses,omitempty"`
	// TeammateTaskSections holds the value of the teammateTaskSections edge.
	TeammateTaskSections []*TeammateTaskSection `json:"teammateTaskSections,omitempty"`
	// TaskLikes holds the value of the taskLikes edge.
	TaskLikes []*TaskLike `json:"taskLikes,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// TeammateTaskColumns holds the value of the teammateTaskColumns edge.
	TeammateTaskColumns []*TeammateTaskColumn `json:"teammateTaskColumns,omitempty"`
	// TeammateTasks holds the value of the teammateTasks edge.
	TeammateTasks []*TeammateTask `json:"teammateTasks,omitempty"`
	// DeletedTasksRef holds the value of the deletedTasksRef edge.
	DeletedTasksRef []*DeletedTask `json:"deletedTasksRef,omitempty"`
	// WorkspaceActivities holds the value of the workspaceActivities edge.
	WorkspaceActivities []*WorkspaceActivity `json:"workspaceActivities,omitempty"`
	// TaskActivities holds the value of the taskActivities edge.
	TaskActivities []*TaskActivity `json:"taskActivities,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [14]bool
}

// TeammateOrErr returns the Teammate value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkspaceEdges) TeammateOrErr() (*Teammate, error) {
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

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[1] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// WorkspaceTeammatesOrErr returns the WorkspaceTeammates value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) WorkspaceTeammatesOrErr() ([]*WorkspaceTeammate, error) {
	if e.loadedTypes[2] {
		return e.WorkspaceTeammates, nil
	}
	return nil, &NotLoadedError{edge: "workspaceTeammates"}
}

// FavoriteWorkspacesOrErr returns the FavoriteWorkspaces value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) FavoriteWorkspacesOrErr() ([]*FavoriteWorkspace, error) {
	if e.loadedTypes[3] {
		return e.FavoriteWorkspaces, nil
	}
	return nil, &NotLoadedError{edge: "favoriteWorkspaces"}
}

// TeammateTaskTabStatusesOrErr returns the TeammateTaskTabStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) TeammateTaskTabStatusesOrErr() ([]*TeammateTaskTabStatus, error) {
	if e.loadedTypes[4] {
		return e.TeammateTaskTabStatuses, nil
	}
	return nil, &NotLoadedError{edge: "teammateTaskTabStatuses"}
}

// TeammateTaskListStatusesOrErr returns the TeammateTaskListStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) TeammateTaskListStatusesOrErr() ([]*TeammateTaskListStatus, error) {
	if e.loadedTypes[5] {
		return e.TeammateTaskListStatuses, nil
	}
	return nil, &NotLoadedError{edge: "teammateTaskListStatuses"}
}

// TeammateTaskSectionsOrErr returns the TeammateTaskSections value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) TeammateTaskSectionsOrErr() ([]*TeammateTaskSection, error) {
	if e.loadedTypes[6] {
		return e.TeammateTaskSections, nil
	}
	return nil, &NotLoadedError{edge: "teammateTaskSections"}
}

// TaskLikesOrErr returns the TaskLikes value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) TaskLikesOrErr() ([]*TaskLike, error) {
	if e.loadedTypes[7] {
		return e.TaskLikes, nil
	}
	return nil, &NotLoadedError{edge: "taskLikes"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[8] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// TeammateTaskColumnsOrErr returns the TeammateTaskColumns value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) TeammateTaskColumnsOrErr() ([]*TeammateTaskColumn, error) {
	if e.loadedTypes[9] {
		return e.TeammateTaskColumns, nil
	}
	return nil, &NotLoadedError{edge: "teammateTaskColumns"}
}

// TeammateTasksOrErr returns the TeammateTasks value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) TeammateTasksOrErr() ([]*TeammateTask, error) {
	if e.loadedTypes[10] {
		return e.TeammateTasks, nil
	}
	return nil, &NotLoadedError{edge: "teammateTasks"}
}

// DeletedTasksRefOrErr returns the DeletedTasksRef value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) DeletedTasksRefOrErr() ([]*DeletedTask, error) {
	if e.loadedTypes[11] {
		return e.DeletedTasksRef, nil
	}
	return nil, &NotLoadedError{edge: "deletedTasksRef"}
}

// WorkspaceActivitiesOrErr returns the WorkspaceActivities value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) WorkspaceActivitiesOrErr() ([]*WorkspaceActivity, error) {
	if e.loadedTypes[12] {
		return e.WorkspaceActivities, nil
	}
	return nil, &NotLoadedError{edge: "workspaceActivities"}
}

// TaskActivitiesOrErr returns the TaskActivities value or an error if the edge
// was not loaded in eager-loading.
func (e WorkspaceEdges) TaskActivitiesOrErr() ([]*TaskActivity, error) {
	if e.loadedTypes[13] {
		return e.TaskActivities, nil
	}
	return nil, &NotLoadedError{edge: "taskActivities"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Workspace) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case workspace.FieldDescription:
			values[i] = new([]byte)
		case workspace.FieldName:
			values[i] = new(sql.NullString)
		case workspace.FieldCreatedAt, workspace.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case workspace.FieldID, workspace.FieldCreatedBy:
			values[i] = new(ulid.ID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Workspace", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Workspace fields.
func (w *Workspace) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workspace.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				w.ID = *value
			}
		case workspace.FieldCreatedBy:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				w.CreatedBy = *value
			}
		case workspace.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				w.Name = value.String
			}
		case workspace.FieldDescription:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &w.Description); err != nil {
					return fmt.Errorf("unmarshal field description: %w", err)
				}
			}
		case workspace.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				w.CreatedAt = value.Time
			}
		case workspace.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				w.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTeammate queries the "teammate" edge of the Workspace entity.
func (w *Workspace) QueryTeammate() *TeammateQuery {
	return (&WorkspaceClient{config: w.config}).QueryTeammate(w)
}

// QueryProjects queries the "projects" edge of the Workspace entity.
func (w *Workspace) QueryProjects() *ProjectQuery {
	return (&WorkspaceClient{config: w.config}).QueryProjects(w)
}

// QueryWorkspaceTeammates queries the "workspaceTeammates" edge of the Workspace entity.
func (w *Workspace) QueryWorkspaceTeammates() *WorkspaceTeammateQuery {
	return (&WorkspaceClient{config: w.config}).QueryWorkspaceTeammates(w)
}

// QueryFavoriteWorkspaces queries the "favoriteWorkspaces" edge of the Workspace entity.
func (w *Workspace) QueryFavoriteWorkspaces() *FavoriteWorkspaceQuery {
	return (&WorkspaceClient{config: w.config}).QueryFavoriteWorkspaces(w)
}

// QueryTeammateTaskTabStatuses queries the "teammateTaskTabStatuses" edge of the Workspace entity.
func (w *Workspace) QueryTeammateTaskTabStatuses() *TeammateTaskTabStatusQuery {
	return (&WorkspaceClient{config: w.config}).QueryTeammateTaskTabStatuses(w)
}

// QueryTeammateTaskListStatuses queries the "teammateTaskListStatuses" edge of the Workspace entity.
func (w *Workspace) QueryTeammateTaskListStatuses() *TeammateTaskListStatusQuery {
	return (&WorkspaceClient{config: w.config}).QueryTeammateTaskListStatuses(w)
}

// QueryTeammateTaskSections queries the "teammateTaskSections" edge of the Workspace entity.
func (w *Workspace) QueryTeammateTaskSections() *TeammateTaskSectionQuery {
	return (&WorkspaceClient{config: w.config}).QueryTeammateTaskSections(w)
}

// QueryTaskLikes queries the "taskLikes" edge of the Workspace entity.
func (w *Workspace) QueryTaskLikes() *TaskLikeQuery {
	return (&WorkspaceClient{config: w.config}).QueryTaskLikes(w)
}

// QueryTags queries the "tags" edge of the Workspace entity.
func (w *Workspace) QueryTags() *TagQuery {
	return (&WorkspaceClient{config: w.config}).QueryTags(w)
}

// QueryTeammateTaskColumns queries the "teammateTaskColumns" edge of the Workspace entity.
func (w *Workspace) QueryTeammateTaskColumns() *TeammateTaskColumnQuery {
	return (&WorkspaceClient{config: w.config}).QueryTeammateTaskColumns(w)
}

// QueryTeammateTasks queries the "teammateTasks" edge of the Workspace entity.
func (w *Workspace) QueryTeammateTasks() *TeammateTaskQuery {
	return (&WorkspaceClient{config: w.config}).QueryTeammateTasks(w)
}

// QueryDeletedTasksRef queries the "deletedTasksRef" edge of the Workspace entity.
func (w *Workspace) QueryDeletedTasksRef() *DeletedTaskQuery {
	return (&WorkspaceClient{config: w.config}).QueryDeletedTasksRef(w)
}

// QueryWorkspaceActivities queries the "workspaceActivities" edge of the Workspace entity.
func (w *Workspace) QueryWorkspaceActivities() *WorkspaceActivityQuery {
	return (&WorkspaceClient{config: w.config}).QueryWorkspaceActivities(w)
}

// QueryTaskActivities queries the "taskActivities" edge of the Workspace entity.
func (w *Workspace) QueryTaskActivities() *TaskActivityQuery {
	return (&WorkspaceClient{config: w.config}).QueryTaskActivities(w)
}

// Update returns a builder for updating this Workspace.
// Note that you need to call Workspace.Unwrap() before calling this method if this Workspace
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Workspace) Update() *WorkspaceUpdateOne {
	return (&WorkspaceClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Workspace entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Workspace) Unwrap() *Workspace {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Workspace is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Workspace) String() string {
	var builder strings.Builder
	builder.WriteString("Workspace(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", w.CreatedBy))
	builder.WriteString(", name=")
	builder.WriteString(w.Name)
	builder.WriteString(", description=")
	builder.WriteString(fmt.Sprintf("%v", w.Description))
	builder.WriteString(", created_at=")
	builder.WriteString(w.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(w.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Workspaces is a parsable slice of Workspace.
type Workspaces []*Workspace

func (w Workspaces) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}

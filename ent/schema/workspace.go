package schema

import (
	"project-management-demo-backend/ent/annotation"
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/contrib/entgql"

	"entgo.io/ent/schema"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

const workspacesRef string = "workspaces"

// Workspace holds the schema definition for the Test entity.
type Workspace struct {
	ent.Schema
}

// WorkspaceMixin defines Fields
type WorkspaceMixin struct {
	entMixin.Schema
}

// Fields of the Workspace.
func (WorkspaceMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("created_by").
			GoType(ulid.ID("")),
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.JSON("description", map[string]interface{}{}).
			Annotations(
				annotation.MutationInput{
					SkipPtr: true,
				},
			),
	}
}

// Edges of the Workspace.
func (Workspace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teammate", Teammate.Type).
			Ref(workspacesRef).
			Unique().
			Field("created_by").
			Required().
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "created_by"},
				),
			),
		edge.To(projectsRef, Project.Type),
		edge.To(workspaceTeammatesRef, WorkspaceTeammate.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_teammate_id"},
				),
			),
		edge.To(favoriteWorkspacesRef, FavoriteWorkspace.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "favorite_workspace_id"},
				),
			),
		edge.To(teammateTaskTabStatusesRef, TeammateTaskTabStatus.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_task_tab_status_id"},
				),
			),
		edge.To(teammateTaskListStatusesRef, TeammateTaskListStatus.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_task_list_status_id"},
				),
			),
		edge.To(teammateTaskSectionsRef, TeammateTaskSection.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_task_section_id"},
				),
			),
		edge.To(taskLikesRef, TaskLike.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_like_id"},
				),
			),
		edge.To(tagsRef, Tag.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "tag_id"},
				),
			),
		edge.To(teammateTaskColumnsRef, TeammateTaskColumn.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_task_column_id"},
				),
			),
		edge.To(teammateTasksRef, TeammateTask.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_task_id"},
				),
			),
		edge.To(deletedTasksRef, DeletedTask.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "deleted_task_id"},
				),
			),
		edge.To(workspaceActivitiesRef, WorkspaceActivity.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_activity_id"},
				),
			),
		edge.To(taskActivitiesRef, TaskActivity.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_activity_id"},
				),
			),
		edge.To(archivedTaskActivitiesRef, ArchivedTaskActivity.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "archived_task_activity_id"},
				),
			),
		edge.To(archivedWorkspaceActivitiesRef, ArchivedWorkspaceActivity.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "archived_workspace_activity_id"},
				),
			),
	}
}

// Mixin of the Workspace.
func (Workspace) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Workspace.Prefix),
		WorkspaceMixin{},
		mixin.NewDatetime(),
	}
}

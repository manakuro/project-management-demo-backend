package schema

import (
	"project-management-demo-backend/ent/annotation"
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/contrib/entgql"

	"entgo.io/ent/schema"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// Teammate holds the schema definition for the Test entity.
type Teammate struct {
	ent.Schema
}

// TeammateMixin defines Fields
type TeammateMixin struct {
	entMixin.Schema
}

// Fields of the Teammate.
func (TeammateMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.String("image").NotEmpty(),
		field.String("email").NotEmpty(),
	}
}

// Edges of the Teammate.
func (Teammate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(workspacesRef, Workspace.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_id"},
				),
			),
		edge.To(projectsRef, Project.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "project_id"},
				),
			),
		edge.To(projectTeammatesRef, ProjectTeammate.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "project_teammate_id"},
				),
			),
		edge.To(workspaceTeammatesRef, WorkspaceTeammate.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_teammate_id"},
				),
			),
		edge.To(favoriteProjectsRef, FavoriteProject.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "favorite_project_id"},
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
		edge.To(teammateTaskColumnsRef, TeammateTaskColumn.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_task_column_id"},
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
		edge.To(tasksRef, Task.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_id"},
				),
			),
		edge.To(teammateTasksRef, TeammateTask.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_task_id"},
				),
			),
		edge.To(taskLikesRef, TaskLike.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_like_id"},
				),
			),
		edge.To(taskCollaboratorsRef, TaskCollaborator.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_collaborator_id"},
				),
			),
		edge.To(taskFeedsRef, TaskFeed.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_feed_id"},
				),
			),
		edge.To(taskFeedLikesRef, TaskFeedLike.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_feed_like_id"},
				),
			),
		edge.To(taskActivitiesRef, TaskActivity.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_activity_id"},
				),
			),
		edge.To(workspaceActivitiesRef, WorkspaceActivity.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_activity_id"},
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

// TODO: Use fulltext search.
//// Indexes of the Teammate.
//func (Teammate) Indexes() []ent.Index {
//	return []ent.Index{
//		index.Fields("email").
//			Annotations(
//				entsql.IndexTypes(map[string]string{
//					dialect.MySQL: "FULLTEXT",
//				})),
//	}
//}

// Mixin of the Teammate.
func (Teammate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Teammate.Prefix),
		TeammateMixin{},
		mixin.NewDatetime(),
	}
}

package schema

import (
	"project-management-demo-backend/ent/annotation"
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/contrib/entgql"

	"entgo.io/ent/dialect"

	"entgo.io/ent/schema"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

const tasksRef string = "tasks"

// Task holds the schema definition for the Test entity.
type Task struct {
	ent.Schema
}

// TaskMixin defines Fields
type TaskMixin struct {
	entMixin.Schema
}

// Fields of the Task.
func (TaskMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("task_parent_id").
			GoType(ulid.ID("")).
			Optional(),
		field.String("task_priority_id").
			GoType(ulid.ID("")).
			Optional(),
		field.String("assignee_id").
			GoType(ulid.ID("")).
			Optional(),
		field.String("created_by").
			GoType(ulid.ID("")).
			Annotations(
				annotation.WhereInput{Type: "ID"},
			),
		field.Bool("completed").
			Default(false),
		field.Time("completed_at").
			Nillable().
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Bool("is_new").
			Default(false),
		field.String("name").
			MaxLen(255),
		field.Time("due_date").
			Nillable().
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("due_time").
			Nillable().
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.JSON("description", map[string]interface{}{}).
			Annotations(
				annotation.MutationInput{
					SkipPtr: true,
				},
			),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teammate", Teammate.Type).
			Ref(tasksRef).
			Unique().
			Field("assignee_id").
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "assignee_id"},
				),
			),
		edge.From("taskPriority", TaskPriority.Type).
			Ref(tasksRef).
			Unique().
			Field("task_priority_id").
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_priority_id"},
				),
			),
		edge.To("subTasks", Task.Type).
			Annotations(entgql.Bind()),

		edge.From("parentTask", Task.Type).
			Ref("subTasks").
			Field("task_parent_id").
			Unique().
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_parent_id"},
				),
			),
		edge.To(teammateTasksRef, TeammateTask.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_task_id"},
				),
			),
		edge.To(projectTasksRef, ProjectTask.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "project_task_id"},
				),
			),
		edge.To(taskLikesRef, TaskLike.Type).Annotations(
			entgql.Bind(),
			schema.Annotation(
				annotation.Edge{FieldName: "task_like_id"},
			),
		),
		edge.To(taskTagsRef, TaskTag.Type).Annotations(
			entgql.Bind(),
			schema.Annotation(
				annotation.Edge{FieldName: "task_tag_id"},
			),
		),
		edge.To(taskCollaboratorsRef, TaskCollaborator.Type).Annotations(
			entgql.Bind(),
			schema.Annotation(
				annotation.Edge{FieldName: "task_collaborator_id"},
			),
		),
		edge.To(taskFeedsRef, TaskFeed.Type).Annotations(
			entgql.Bind(),
			schema.Annotation(
				annotation.Edge{FieldName: "task_feed_id"},
			),
		),
		edge.To(taskFeedLikesRef, TaskFeedLike.Type).Annotations(
			entgql.Bind(),
			schema.Annotation(
				annotation.Edge{FieldName: "task_feed_like_id"},
			),
		),
		edge.To(taskFilesRef, TaskFile.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_file_id"},
				),
			),
		edge.To(deletedTasksRef, DeletedTask.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "deleted_task_id"},
				),
			),
		edge.To(taskActivityTasksRef, TaskActivityTask.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_activity_task_id"},
				),
			),
		edge.To(workspaceActivityTasksRef, WorkspaceActivityTask.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_activity_task_id"},
				),
			),
		edge.To(archivedTaskActivityTasksRef, ArchivedTaskActivityTask.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "archived_activity_task_id"},
				),
			),
		edge.To(archivedWorkspaceActivityTasksRef, ArchivedWorkspaceActivityTask.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "archived_workspace_activity_task_id"},
				),
			),
		edge.To(deletedTeammateTasksRef, DeletedTeammateTask.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "deleted_teammate_task_id"},
				),
			),
	}
}

// Annotations of the ProjectTask.
func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Annotation(
			annotation.MutationInput{
				Create: []annotation.MutationInputField{
					{
						Key:  "WorkspaceID",
						Type: "ulid.ID",
					},
				},
				Update: []annotation.MutationInputField{
					{
						Key:  "WorkspaceID",
						Type: "ulid.ID",
					},
				},
			},
		),
	}
}

// Mixin of the Task.
func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Task.Prefix),
		TaskMixin{},
		mixin.NewDatetime(),
	}
}

package schema

import (
	"project-management-demo-backend/ent/annotation"
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/ent/schema"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// ProjectTaskListStatus holds the schema definition for the Test entity.
type ProjectTaskListStatus struct {
	ent.Schema
}

// ProjectTaskListStatusMixin defines Fields
type ProjectTaskListStatusMixin struct {
	entMixin.Schema
}

// Fields of the ProjectTaskListStatus.
func (ProjectTaskListStatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("project_id").
			GoType(ulid.ID("")),
		field.String("task_list_completed_status_id").
			GoType(ulid.ID("")),
		field.String("task_list_sort_status_id").
			GoType(ulid.ID("")),
	}
}

// Edges of the ProjectTaskListStatus.
func (ProjectTaskListStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("project_task_list_statuses").
			Field("project_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "project_id"},
				),
			),
		edge.From("task_list_completed_status", TaskListCompletedStatus.Type).
			Ref("project_task_list_statuses").
			Field("task_list_completed_status_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "task_list_completed_status_id"},
				),
			),
		edge.From("task_list_sort_status", TaskListSortStatus.Type).
			Ref("project_task_list_statuses").
			Field("task_list_sort_status_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "task_list_sort_status_id"},
				),
			),
	}
}

// Mixin of the ProjectTaskListStatus.
func (ProjectTaskListStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().ProjectTaskListStatus.Prefix),
		ProjectTaskListStatusMixin{},
		mixin.NewDatetime(),
	}
}

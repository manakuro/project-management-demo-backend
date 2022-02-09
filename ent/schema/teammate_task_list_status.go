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

const teammateTaskListStatusesRef string = "teammateTaskListStatuses"

// TeammateTaskListStatus holds the schema definition for the Test entity.
type TeammateTaskListStatus struct {
	ent.Schema
}

// TeammateTaskListStatusMixin defines Fields
type TeammateTaskListStatusMixin struct {
	entMixin.Schema
}

// Fields of the TeammateTaskListStatus.
func (TeammateTaskListStatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("workspace_id").
			GoType(ulid.ID("")),
		field.String("teammate_id").
			GoType(ulid.ID("")),
		field.String("task_list_completed_status_id").
			GoType(ulid.ID("")),
		field.String("task_list_sort_status_id").
			GoType(ulid.ID("")),
	}
}

// Edges of the TeammateTaskListStatus.
func (TeammateTaskListStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("workspace", Workspace.Type).
			Ref(teammateTaskListStatusesRef).
			Field("workspace_id").
			Unique().
			Required().
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_id"},
				),
			),
		edge.From("teammate", Teammate.Type).
			Ref(teammateTaskListStatusesRef).
			Field("teammate_id").
			Unique().
			Required().
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_id"},
				),
			),
		edge.From("task_list_completed_status", TaskListCompletedStatus.Type).
			Ref(teammateTaskListStatusesRef).
			Field("task_list_completed_status_id").
			Unique().
			Required().
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_list_completed_status_id"},
				),
			),
		edge.From("task_list_sort_status", TaskListSortStatus.Type).
			Ref(teammateTaskListStatusesRef).
			Field("task_list_sort_status_id").
			Unique().
			Required().
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_list_sort_status_id"},
				),
			),
	}
}

// Mixin of the TeammateTaskListStatus.
func (TeammateTaskListStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().TeammateTaskListStatus.Prefix),
		TeammateTaskListStatusMixin{},
		mixin.NewDatetime(),
	}
}

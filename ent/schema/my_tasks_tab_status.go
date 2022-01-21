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

// MyTasksTabStatus holds the schema definition for the Test entity.
type MyTasksTabStatus struct {
	ent.Schema
}

// MyTasksTabStatusMixin defines Fields
type MyTasksTabStatusMixin struct {
	entMixin.Schema
}

// Fields of the MyTasksTabStatus.
func (MyTasksTabStatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("workspace_id").
			GoType(ulid.ID("")),
		field.String("teammate_id").
			GoType(ulid.ID("")),
		field.Enum("status").
			NamedValues(
				"List", "LIST",
				"Board", "BOARD",
				"Calendar", "CALENDAR",
				"Files", "FILES",
			).Default("LIST"),
	}
}

// Edges of the MyTasksTabStatus.
func (MyTasksTabStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("workspace", Workspace.Type).
			Ref("my_tasks_tab_statuses").
			Field("workspace_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_id"},
				),
			),
		edge.From("teammate", Teammate.Type).
			Ref("my_tasks_tab_statuses").
			Field("teammate_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_id"},
				),
			),
	}
}

// Mixin of the MyTasksTabStatus.
func (MyTasksTabStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().MyTasksTabStatus.Prefix),
		MyTasksTabStatusMixin{},
		mixin.NewDatetime(),
	}
}

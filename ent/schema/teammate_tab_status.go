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

// TeammateTabStatus holds the schema definition for the Test entity.
type TeammateTabStatus struct {
	ent.Schema
}

// TeammateTabStatusMixin defines Fields
type TeammateTabStatusMixin struct {
	entMixin.Schema
}

// Fields of the TeammateTabStatus.
func (TeammateTabStatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("workspace_id").
			GoType(ulid.ID("")),
		field.String("teammate_id").
			GoType(ulid.ID("")),
		field.Enum("status_code").
			NamedValues(
				"List", "LIST",
				"Board", "BOARD",
				"Calendar", "CALENDAR",
				"Files", "FILES",
			).Default("LIST"),
	}
}

// Edges of the TeammateTabStatus.
func (TeammateTabStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("workspace", Workspace.Type).
			Ref("teammate_tab_statuses").
			Field("workspace_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_id"},
				),
			),
		edge.From("teammate", Teammate.Type).
			Ref("teammate_tab_statuses").
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

// Mixin of the TeammateTabStatus.
func (TeammateTabStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().TeammateTabStatus.Prefix),
		TeammateTabStatusMixin{},
		mixin.NewDatetime(),
	}
}

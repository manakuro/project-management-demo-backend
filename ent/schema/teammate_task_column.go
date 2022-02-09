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

const teammateTaskColumnsRef string = "teammateTaskColumns"

// TeammateTaskColumn holds the schema definition for the Test entity.
type TeammateTaskColumn struct {
	ent.Schema
}

// TeammateTaskColumnMixin defines Fields
type TeammateTaskColumnMixin struct {
	entMixin.Schema
}

// Fields of the TeammateTaskColumn.
func (TeammateTaskColumnMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("teammate_id").
			GoType(ulid.ID("")),
		field.String("task_column_id").
			GoType(ulid.ID("")),
		field.String("workspace_id").
			GoType(ulid.ID("")),
		field.String("width").
			NotEmpty().
			MaxLen(255),
		field.Bool("disabled"),
		field.Bool("customizable"),
		field.Int("order"),
	}
}

// Edges of the TeammateTaskColumn.
func (TeammateTaskColumn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teammate", Teammate.Type).
			Ref(teammateTaskColumnsRef).
			Field("teammate_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_id"},
				),
			),
		edge.From("workspace", Workspace.Type).
			Ref(teammateTaskColumnsRef).
			Field("workspace_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_id"},
				),
			),
		edge.From("task_column", TaskColumn.Type).
			Ref(teammateTaskColumnsRef).
			Field("task_column_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "task_column_id"},
				),
			),
	}
}

// Mixin of the TeammateTaskColumn.
func (TeammateTaskColumn) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().TeammateTaskColumn.Prefix),
		TeammateTaskColumnMixin{},
		mixin.NewDatetime(),
	}
}

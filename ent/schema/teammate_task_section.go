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

// TeammateTaskSection holds the schema definition for the Test entity.
type TeammateTaskSection struct {
	ent.Schema
}

// TeammateTaskSectionMixin defines Fields
type TeammateTaskSectionMixin struct {
	entMixin.Schema
}

// Fields of the TeammateTaskSection.
func (TeammateTaskSectionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("teammate_id").
			GoType(ulid.ID("")),
		field.String("workspace_id").
			GoType(ulid.ID("")),
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.Bool("assigned"),
	}
}

// Edges of the TeammateTaskSection.
func (TeammateTaskSection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teammate", Teammate.Type).
			Ref("teammate_task_sections").
			Field("teammate_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_id"},
				),
			),
		edge.From("workspace", Workspace.Type).
			Ref("teammate_task_sections").
			Field("workspace_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_id"},
				),
			),
		edge.To("teammate_tasks", TeammateTask.Type).Annotations(
			schema.Annotation(
				annotation.Edge{FieldName: "teammate_task_id"},
			),
		),
	}
}

// Mixin of the TeammateTaskSection.
func (TeammateTaskSection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().TeammateTaskSection.Prefix),
		TeammateTaskSectionMixin{},
		mixin.NewDatetime(),
	}
}
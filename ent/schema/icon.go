package schema

import (
	"project-management-demo-backend/ent/annotation"
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// Icon holds the schema definition for the Test entity.
type Icon struct {
	ent.Schema
}

// IconMixin defines Fields
type IconMixin struct {
	entMixin.Schema
}

// Fields of the Icon.
func (IconMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.String("icon").
			NotEmpty().
			MaxLen(255),
	}
}

// Edges of the Icon.
func (Icon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("project", Project.Type).
			Unique().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "project_id"},
				),
			),
	}
}

// Mixin of the Icon.
func (Icon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Icon.Prefix),
		IconMixin{},
		mixin.NewDatetime(),
	}
}

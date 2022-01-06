package schema

import (
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// Color holds the schema definition for the Test entity.
type Color struct {
	ent.Schema
}

// ColorMixin defines Fields
type ColorMixin struct {
	entMixin.Schema
}

// Fields of the Color.
func (ColorMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.String("color").
			NotEmpty().
			MaxLen(255),
		field.String("hex").
			NotEmpty().
			MaxLen(255),
	}
}

// Edges of the Color.
func (Color) Edges() []ent.Edge {
	return nil
}

// Mixin of the Color.
func (Color) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Color.Prefix),
		ColorMixin{},
		mixin.NewDatetime(),
	}
}

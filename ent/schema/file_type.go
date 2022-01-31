package schema

import (
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// FileType holds the schema definition for the Test entity.
type FileType struct {
	ent.Schema
}

// FileTypeMixin defines Fields
type FileTypeMixin struct {
	entMixin.Schema
}

// Fields of the FileType.
func (FileTypeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.Enum("type_code").
			NamedValues(
				"Image", "IMAGE",
				"PDF", "PDF",
				"Text", "TEXT",
			),
	}
}

// Edges of the FileType.
func (FileType) Edges() []ent.Edge {
	return nil
}

// Mixin of the FileType.
func (FileType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().FileType.Prefix),
		FileTypeMixin{},
		mixin.NewDatetime(),
	}
}

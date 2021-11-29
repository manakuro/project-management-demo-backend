package schema

import (
	"project-management-demo-backend/ent/mixin"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// TestUser holds the schema definition for the Test entity.
type TestUser struct {
	ent.Schema
}

var testUserTablePrefix = "TU"

// TestUserMixin defines Fields
type TestUserMixin struct {
	entMixin.Schema
}

// Fields of the TestUser.
func (TestUserMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Int("age"),
	}
}

// Edges of the TestUser.
func (TestUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("test_todos", TestTodo.Type),
	}
}

// Mixin of the TestUser.
func (TestUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(testUserTablePrefix),
		TestUserMixin{},
		mixin.NewDatetime(),
	}
}

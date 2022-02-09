package schema

import (
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/ent/schema/testuserprofile"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/contrib/entgql"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// TestUser holds the schema definition for the Test entity.
type TestUser struct {
	ent.Schema
}

// TestUserMixin defines Fields
type TestUserMixin struct {
	entMixin.Schema
}

// Fields of the TestUser.
func (TestUserMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.Int("age"),
		field.JSON("profile", testuserprofile.TestUserProfile{}),
	}
}

// Edges of the TestUser.
func (TestUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("testTodos", TestTodo.Type).
			Annotations(entgql.Bind()),
	}
}

// Mixin of the TestUser.
func (TestUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().TestUser.Prefix),
		TestUserMixin{},
		mixin.NewDatetime(),
	}
}

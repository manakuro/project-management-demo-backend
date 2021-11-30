package schema

import (
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/ent/schema/edge"
	entMixin "entgo.io/ent/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TestTodo holds the schema definition for the TestTodo entity.
type TestTodo struct {
	ent.Schema
}

// TestTodoMixin defines Fields
type TestTodoMixin struct {
	entMixin.Schema
}

// Fields of the TestTodo.
func (TestTodoMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("test_user_id").
			GoType(ulid.ID("")).
			Optional(),
		field.String("name").Default(""),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS"),
		field.Int("priority").Default(0),
	}
}

// Edges of the TestTodo.
func (TestTodo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("test_user", TestUser.Type).
			Ref("test_todos").
			Unique().
			Field("test_user_id"),
	}
}

// Mixin of the TestTodo.
func (TestTodo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().TestTodo.Prefix),
		TestTodoMixin{},
		mixin.NewDatetime(),
	}
}

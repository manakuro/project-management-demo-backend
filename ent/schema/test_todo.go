package schema

import (
	"project-management-demo-backend/ent/schema/pulid"
	"time"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// TestTodo holds the schema definition for the TestTodo entity.
type TestTodo struct {
	ent.Schema
}

// Fields of the TestTodo.
func (TestTodo) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(pulid.ID("")).
			DefaultFunc(func() pulid.ID { return pulid.MustNew("") }),
		field.String("name").NotEmpty(),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS"),
		field.Int("priority").Default(0),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime DEFAULT CURRENT_TIMESTAMP",
			}),
		field.Time("updated_at").
			Immutable().
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
			}),
	}
}

// Edges of the TestTodo.
func (TestTodo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("test_user", TestUser.Type).
			Ref("test_todos").
			Unique(),
	}
}

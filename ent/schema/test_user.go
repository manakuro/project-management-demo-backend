package schema

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// TestUser holds the schema definition for the Test entity.
type TestUser struct {
	ent.Schema
}

// Fields of the Test.
func (TestUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(ulid.ID("")).
			DefaultFunc(func() ulid.ID {
				return ulid.MustNew("")
			}),
		field.String("name").
			NotEmpty(),
		field.Int("age"),
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime DEFAULT CURRENT_TIMESTAMP",
			}),
		field.Time("updated_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
			}),
	}
}

// Edges of the Test.
func (TestUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("test_todos", TestTodo.Type),
	}
}

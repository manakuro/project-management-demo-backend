package schema

import (
	"time"

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
		field.String("name").NotEmpty(),
		field.Enum("status").
			Values("in_progress", "completed").
			Default("in_progress"),
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
				dialect.MySQL: "datetime DEFAULT CURRENT_TIMESTAMP",
			}),
	}
}

// Edges of the TestTodo.
func (TestTodo) Edges() []ent.Edge {
	return nil
}

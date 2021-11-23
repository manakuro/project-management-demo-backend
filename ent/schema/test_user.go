package schema

import (
	"time"

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
		field.String("name").Default(""),
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
	return nil
}

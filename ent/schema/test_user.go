package schema

import (
	"project-management-demo-backend/ent/schema/ulid"
	"time"

	"entgo.io/ent/schema"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// TestUser holds the schema definition for the Test entity.
type TestUser struct {
	ent.Schema
}

var testUserTablePrefix = "TU"

// Fields of the Test.
func (TestUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(ulid.ID("")).
			DefaultFunc(func() ulid.ID {
				return ulid.MustNew(testUserTablePrefix)
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

// Annotation captures the id prefix for a type.
type Annotation struct {
	Prefix string
}

// Name implements the ent Annotation interface.
func (a Annotation) Name() string {
	return "ULID"
}

// Annotations of the TestUser
func (TestUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{Prefix: testUserTablePrefix},
	}
}

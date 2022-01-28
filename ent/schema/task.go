package schema

import (
	"project-management-demo-backend/ent/annotation"
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/ent/dialect"

	"entgo.io/ent/schema"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// Task holds the schema definition for the Test entity.
type Task struct {
	ent.Schema
}

// TaskMixin defines Fields
type TaskMixin struct {
	entMixin.Schema
}

// Fields of the Task.
func (TaskMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("task_parent_id").
			GoType(ulid.ID("")).
			Optional(),
		field.String("task_priority_id").
			GoType(ulid.ID("")),
		field.String("assignee_id").
			GoType(ulid.ID("")).
			Optional(),
		field.String("created_by").
			GoType(ulid.ID("")).
			Annotations(
				annotation.WhereInput{Type: "ID"},
			),
		field.Bool("completed").
			Default(false),
		field.Time("completed_at").
			Nillable().
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Bool("is_new").
			Default(false),
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.Time("due_date").
			Nillable().
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("due_time").
			Nillable().
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teammate", Teammate.Type).
			Ref("tasks").
			Unique().
			Field("assignee_id").
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "assignee_id"},
				),
			),
		edge.From("task_priority", TaskPriority.Type).
			Ref("tasks").
			Unique().
			Field("task_priority_id").
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "task_priority_id"},
				),
			),
		edge.To("sub_tasks", Task.Type).
			From("parent").
			Field("task_parent_id").
			Unique().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "task_parent_id"},
				),
			),
		edge.To("teammate_tasks", TeammateTask.Type).Annotations(
			schema.Annotation(
				annotation.Edge{FieldName: "teammate_task_id"},
			),
		),
		edge.To("project_tasks", ProjectTask.Type).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "project_task_id"},
				),
			),
	}
}

// Mixin of the Task.
func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Task.Prefix),
		TaskMixin{},
		mixin.NewDatetime(),
	}
}

package schema

import (
	"project-management-demo-backend/ent/annotation"
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/ent/schema/editor"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/ent/schema"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// TaskFeed holds the schema definition for the Test entity.
type TaskFeed struct {
	ent.Schema
}

// TaskFeedMixin defines Fields
type TaskFeedMixin struct {
	entMixin.Schema
}

// Fields of the TaskFeed.
func (TaskFeedMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("task_id").
			GoType(ulid.ID("")),
		field.String("teammate_id").
			GoType(ulid.ID("")),
		field.JSON("description", editor.Description{}),
		field.Bool("is_first").
			Default(false),
		field.Bool("is_pinned").
			Default(false),
	}
}

// Edges of the TaskFeed.
func (TaskFeed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).
			Ref("task_feeds").
			Field("task_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "task_id"},
				),
			),
		edge.From("teammate", Teammate.Type).
			Ref("task_feeds").
			Field("teammate_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_id"},
				),
			),
	}
}

// Mixin of the TaskFeed.
func (TaskFeed) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().TaskFeed.Prefix),
		TaskFeedMixin{},
		mixin.NewDatetime(),
	}
}

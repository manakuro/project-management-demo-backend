package schema

import (
	"project-management-demo-backend/ent/annotation"
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/ent/schema"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// TaskFeedLike holds the schema definition for the Test entity.
type TaskFeedLike struct {
	ent.Schema
}

// TaskFeedLikeMixin defines Fields
type TaskFeedLikeMixin struct {
	entMixin.Schema
}

// Fields of the TaskFeedLike.
func (TaskFeedLikeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("task_id").
			GoType(ulid.ID("")),
		field.String("teammate_id").
			GoType(ulid.ID("")),
		field.String("task_feed_id").
			GoType(ulid.ID("")),
	}
}

// Edges of the TaskFeedLike.
func (TaskFeedLike) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).
			Ref("task_feed_likes").
			Field("task_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "task_id"},
				),
			),
		edge.From("teammate", Teammate.Type).
			Ref("task_feed_likes").
			Field("teammate_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_id"},
				),
			),
		edge.From("feed", TaskFeed.Type).
			Ref("task_feed_likes").
			Field("task_feed_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "task_feed_id"},
				),
			),
	}
}

// Mixin of the TaskFeedLike.
func (TaskFeedLike) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().TaskFeedLike.Prefix),
		TaskFeedLikeMixin{},
		mixin.NewDatetime(),
	}
}

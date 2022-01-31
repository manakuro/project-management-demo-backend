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

// TaskTag holds the schema definition for the Test entity.
type TaskTag struct {
	ent.Schema
}

// TaskTagMixin defines Fields
type TaskTagMixin struct {
	entMixin.Schema
}

// Fields of the TaskTag.
func (TaskTagMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("task_id").
			GoType(ulid.ID("")),
		field.String("tag_id").
			GoType(ulid.ID("")),
	}
}

// Edges of the TaskTag.
func (TaskTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).
			Ref("task_tags").
			Field("task_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "task_id"},
				),
			),
		edge.From("tag", Tag.Type).
			Ref("task_tags").
			Field("tag_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "tag_id"},
				),
			),
	}
}

// Mixin of the TaskTag.
func (TaskTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().TaskTag.Prefix),
		TaskTagMixin{},
		mixin.NewDatetime(),
	}
}
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

// TaskCollaborator holds the schema definition for the Test entity.
type TaskCollaborator struct {
	ent.Schema
}

// TaskCollaboratorMixin defines Fields
type TaskCollaboratorMixin struct {
	entMixin.Schema
}

// Fields of the TaskCollaborator.
func (TaskCollaboratorMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("task_id").
			GoType(ulid.ID("")),
		field.String("teammate_id").
			GoType(ulid.ID("")),
	}
}

// Edges of the TaskCollaborator.
func (TaskCollaborator) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).
			Ref("task_collaborators").
			Field("task_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "task_id"},
				),
			),
		edge.From("teammate", Teammate.Type).
			Ref("task_collaborators").
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

// Mixin of the TaskCollaborator.
func (TaskCollaborator) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().TaskCollaborator.Prefix),
		TaskCollaboratorMixin{},
		mixin.NewDatetime(),
	}
}

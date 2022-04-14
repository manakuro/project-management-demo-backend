package schema

import (
	"project-management-demo-backend/ent/annotation"
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// ActivityType holds the schema definition for the Test entity.
type ActivityType struct {
	ent.Schema
}

// ActivityTypeMixin defines Fields
type ActivityTypeMixin struct {
	entMixin.Schema
}

// Fields of the ActivityType.
func (ActivityTypeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.Enum("type_code").
			NamedValues(
				"Task", "TASK",
				"Workspace", "WORKSPACE",
			),
	}
}

// Edges of the ActivityType.
func (ActivityType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(taskActivitiesRef, TaskActivity.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_activity_id"},
				),
			),
	}
}

// Mixin of the ActivityType.
func (ActivityType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().ActivityType.Prefix),
		ActivityTypeMixin{},
		mixin.NewDatetime(),
	}
}

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

// ProjectTaskColumn holds the schema definition for the Test entity.
type ProjectTaskColumn struct {
	ent.Schema
}

// ProjectTaskColumnMixin defines Fields
type ProjectTaskColumnMixin struct {
	entMixin.Schema
}

// Fields of the ProjectTaskColumn.
func (ProjectTaskColumnMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("project_id").
			GoType(ulid.ID("")),
		field.String("task_column_id").
			GoType(ulid.ID("")),
		field.String("width").
			NotEmpty().
			MaxLen(255),
		field.Bool("disabled"),
		field.Bool("customizable"),
		field.Int("order"),
	}
}

// Edges of the ProjectTaskColumn.
func (ProjectTaskColumn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("project_task_columns").
			Field("project_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "project_id"},
				),
			),
		edge.From("task_column", TaskColumn.Type).
			Ref("project_task_columns").
			Field("task_column_id").
			Unique().
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "task_column_id"},
				),
			),
	}
}

// Mixin of the ProjectTaskColumn.
func (ProjectTaskColumn) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().ProjectTaskColumn.Prefix),
		ProjectTaskColumnMixin{},
		mixin.NewDatetime(),
	}
}

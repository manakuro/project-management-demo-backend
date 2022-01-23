package schema

import (
	"project-management-demo-backend/ent/annotation"
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/ent/schema/editor"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/pkg/const/globalid"
	"time"

	"entgo.io/ent/dialect"

	"entgo.io/ent/schema"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// Project holds the schema definition for the Test entity.
type Project struct {
	ent.Schema
}

// ProjectMixin defines Fields
type ProjectMixin struct {
	entMixin.Schema
}

// Fields of the Project.
func (ProjectMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("workspace_id").
			GoType(ulid.ID("")),
		field.String("project_base_color_id").
			GoType(ulid.ID("")),
		field.String("project_light_color_id").
			GoType(ulid.ID("")),
		field.String("project_icon_id").
			GoType(ulid.ID("")),
		field.String("created_by").
			GoType(ulid.ID("")),
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.JSON("description", editor.Description{}),
		field.String("description_title").
			MaxLen(255),
		field.Time("due_date").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("workspace", Workspace.Type).
			Ref("projects").
			Unique().
			Field("workspace_id").
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_id"},
				),
			),
		edge.From("project_base_color", ProjectBaseColor.Type).
			Ref("projects").
			Unique().
			Field("project_base_color_id").
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "project_base_color_id"},
				),
			),
		edge.From("project_light_color", ProjectLightColor.Type).
			Ref("projects").
			Unique().
			Field("project_light_color_id").
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "project_light_color_id"},
				),
			),
		edge.From("project_icon", ProjectIcon.Type).
			Ref("projects").
			Unique().
			Field("project_icon_id").
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "project_icon_id"},
				),
			),
		edge.From("teammate", Teammate.Type).
			Ref("projects").
			Unique().
			Field("created_by").
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "created_by"},
				),
			),
		edge.To("project_teammates", ProjectTeammate.Edges).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "project_teammate_id"},
				),
			),
		edge.To("favorite_projects", FavoriteProject.Type).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "favorite_project_id"},
				),
			),
		edge.To("project_task_columns", ProjectTaskColumn.Type).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "project_task_column_id"},
				),
			),
	}
}

// Mixin of the Project.
func (Project) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Project.Prefix),
		ProjectMixin{},
		mixin.NewDatetime(),
	}
}

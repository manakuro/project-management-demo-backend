package schema

import (
	"project-management-demo-backend/ent/annotation"
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/ent/schema"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// Teammate holds the schema definition for the Test entity.
type Teammate struct {
	ent.Schema
}

// TeammateMixin defines Fields
type TeammateMixin struct {
	entMixin.Schema
}

// Fields of the Teammate.
func (TeammateMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.String("image").NotEmpty(),
		field.String("email").NotEmpty(),
	}
}

// Edges of the Teammate.
func (Teammate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("workspaces", Workspace.Type).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_id"},
				),
			),
		edge.To("projects", Project.Type).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "project_id"},
				),
			),
		edge.To("project_teammates", ProjectTeammate.Type).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "project_teammate_id"},
				),
			),
		edge.To("workspace_teammates", WorkspaceTeammate.Type).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_teammate_id"},
				),
			),
		edge.To("favorite_projects", FavoriteProject.Type).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "favorite_project_id"},
				),
			),
	}
}

// Mixin of the Teammate.
func (Teammate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Teammate.Prefix),
		TeammateMixin{},
		mixin.NewDatetime(),
	}
}

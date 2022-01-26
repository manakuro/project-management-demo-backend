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

// Workspace holds the schema definition for the Test entity.
type Workspace struct {
	ent.Schema
}

// WorkspaceMixin defines Fields
type WorkspaceMixin struct {
	entMixin.Schema
}

// Fields of the Workspace.
func (WorkspaceMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("created_by").
			GoType(ulid.ID("")),
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.JSON("description", editor.Description{}),
	}
}

// Edges of the Workspace.
func (Workspace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teammate", Teammate.Type).
			Ref("workspaces").
			Unique().
			Field("created_by").
			Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "created_by"},
				),
			),
		edge.To("projects", Project.Type),
		edge.To("workspace_teammates", WorkspaceTeammate.Type).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_teammate_id"},
				),
			),
		edge.To("favorite_workspaces", FavoriteWorkspace.Type).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "favorite_workspace_id"},
				),
			),
		edge.To("teammate_task_tab_statuses", TeammateTaskTabStatus.Type).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_task_tab_status_id"},
				),
			),
	}
}

// Mixin of the Workspace.
func (Workspace) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().Workspace.Prefix),
		WorkspaceMixin{},
		mixin.NewDatetime(),
	}
}

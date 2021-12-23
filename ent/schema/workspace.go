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

// Workspace holds the schema definition for the Test entity.
type Workspace struct {
	ent.Schema
}

// WorkspaceMixin defines Fields
type WorkspaceMixin struct {
	entMixin.Schema
}

// WorkspaceAttrs is a type of attrs
type WorkspaceAttrs struct {
	MentionID   string `json:"mentionId"`
	MentionType string `json:"mentionType"`
}

// WorkspaceContent is a type of content
type WorkspaceContent struct {
	Type  string         `json:"type"`
	Text  string         `json:"text"`
	Attrs WorkspaceAttrs `json:"attrs"`
}

// WorkspaceDescription is a type of description json
type WorkspaceDescription struct {
	Type    string `json:"type"`
	Content WorkspaceContent
}

// Fields of the Workspace.
func (WorkspaceMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("created_by").
			GoType(ulid.ID("")),
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.JSON("description", WorkspaceDescription{}),
	}
}

// Edges of the Workspace.
func (Workspace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teammate", Teammate.Type).
			Ref("workspaces").
			Unique().
			Field("created_by").Required().
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "created_by"},
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

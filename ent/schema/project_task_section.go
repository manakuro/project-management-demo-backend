package schema

import (
	"project-management-demo-backend/ent/annotation"
	"project-management-demo-backend/ent/mixin"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/pkg/const/globalid"

	"entgo.io/contrib/entgql"

	"entgo.io/ent/schema"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

const projectTaskSectionsRef string = "projectTaskSections"

// ProjectTaskSection holds the schema definition for the Test entity.
type ProjectTaskSection struct {
	ent.Schema
}

// ProjectTaskSectionMixin defines Fields
type ProjectTaskSectionMixin struct {
	entMixin.Schema
}

// Fields of the ProjectTaskSection.
func (ProjectTaskSectionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("project_id").
			GoType(ulid.ID("")),
		field.String("name").
			MaxLen(255),
	}
}

// Edges of the ProjectTaskSection.
func (ProjectTaskSection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref(projectTaskSectionsRef).
			Field("project_id").
			Unique().
			Required().
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "project_id"},
				),
			),
		edge.To(projectTasksRef, ProjectTask.Type).
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "project_task_id"},
				),
			),
	}
}

// Annotations of the ProjectTask.
func (ProjectTaskSection) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Annotation(
			annotation.MutationInput{
				Create: []annotation.MutationInputField{
					{
						Key:  "WorkspaceID",
						Type: "ulid.ID",
					},
				},
				Update: []annotation.MutationInputField{
					{
						Key:  "WorkspaceID",
						Type: "ulid.ID",
					},
				},
			},
		),
	}
}

// Mixin of the ProjectTaskSection.
func (ProjectTaskSection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().ProjectTaskSection.Prefix),
		ProjectTaskSectionMixin{},
		mixin.NewDatetime(),
	}
}

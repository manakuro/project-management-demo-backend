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

const teammateTasksRef string = "teammateTasks"

// TeammateTask holds the schema definition for the Test entity.
type TeammateTask struct {
	ent.Schema
}

// TeammateTaskMixin defines Fields
type TeammateTaskMixin struct {
	entMixin.Schema
}

// Fields of the TeammateTask.
func (TeammateTaskMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("teammate_id").
			GoType(ulid.ID("")),
		field.String("task_id").
			GoType(ulid.ID("")),
		field.String("teammate_task_section_id").
			GoType(ulid.ID("")),
		field.String("workspace_id").
			GoType(ulid.ID("")),
	}
}

// Edges of the TeammateTask.
func (TeammateTask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teammate", Teammate.Type).
			Ref(teammateTasksRef).
			Field("teammate_id").
			Unique().
			Required().
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_id"},
				),
			),
		edge.From("task", Task.Type).
			Ref(teammateTasksRef).
			Field("task_id").
			Unique().
			Required().
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "task_id"},
				),
			),
		edge.From("teammateTaskSection", TeammateTaskSection.Type).
			Ref(teammateTasksRef).
			Field("teammate_task_section_id").
			Unique().
			Required().
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_task_section_id"},
				),
			),
		edge.From("workspace", Workspace.Type).
			Ref(teammateTasksRef).
			Field("workspace_id").
			Unique().
			Required().
			Annotations(
				entgql.Bind(),
				schema.Annotation(
					annotation.Edge{FieldName: "workspace_id"},
				),
			),
	}
}

// Annotations of the TeammateTask.
func (TeammateTask) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Annotation(
			annotation.MutationInput{
				Create: []annotation.MutationInputField{
					{
						Key:  "TaskParentID",
						Type: "*ulid.ID",
					},
				},
			},
		),
	}
}

// Mixin of the TeammateTask.
func (TeammateTask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().TeammateTask.Prefix),
		TeammateTaskMixin{},
		mixin.NewDatetime(),
	}
}

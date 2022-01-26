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

// TaskListCompletedStatus holds the schema definition for the Test entity.
type TaskListCompletedStatus struct {
	ent.Schema
}

// TaskListCompletedStatusMixin defines Fields
type TaskListCompletedStatusMixin struct {
	entMixin.Schema
}

// Fields of the TaskListCompletedStatus.make
func (TaskListCompletedStatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.Enum("status_code").
			NamedValues(
				"Incomplete", "INCOMPLETE",
				"Completed", "COMPLETED",
				"CompletedToday", "COMPLETED_TODAY",
				"CompletedYesterday", "COMPLETED_YESTERDAY",
				"Completed1Week", "COMPLETED_1_WEEK",
				"Completed2Weeks", "COMPLETED_2_WEEKS",
				"Completed3Weeks", "COMPLETED_3_WEEKS",
				"All", "ALL",
			),
	}
}

// Edges of the TaskListCompletedStatus.
func (TaskListCompletedStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("teammate_task_list_statuses", TeammateTaskListStatus.Type).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "teammate_task_list_status_id"},
				),
			),
		edge.To("project_task_list_statuses", ProjectTaskListStatus.Type).
			Annotations(
				schema.Annotation(
					annotation.Edge{FieldName: "project_task_list_status_id"},
				),
			),
	}
}

// Mixin of the TaskListCompletedStatus.
func (TaskListCompletedStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.NewUlid(globalid.New().TaskListCompletedStatus.Prefix),
		TaskListCompletedStatusMixin{},
		mixin.NewDatetime(),
	}
}

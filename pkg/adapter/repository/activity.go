package repository

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/activitytype"
	"project-management-demo-backend/ent/taskactivity"
	"project-management-demo-backend/ent/workspaceactivity"
	"project-management-demo-backend/pkg/entity/model"
	ur "project-management-demo-backend/pkg/usecase/repository"

	"entgo.io/ent/dialect/sql"
)

type activityRepository struct {
	client *ent.Client
}

// NewActivityRepository generates teammate repository
func NewActivityRepository(client *ent.Client) ur.Activity {
	return &activityRepository{client: client}
}

func (r *activityRepository) List(ctx context.Context, where model.ActivityWhereInput) ([]*model.Activity, error) {
	db := r.client.DB()

	taskActivityTable := sql.Table(taskactivity.Table).As(taskactivity.Table)
	workspaceActivityTable := sql.Table(workspaceactivity.Table).As(workspaceactivity.Table)
	activityTypeTable := sql.Table(activitytype.Table).As(activitytype.Table)

	taskActivitySelector := sql.Select(
		sql.As(taskActivityTable.C(taskactivity.FieldID), "id"),
		sql.As(activityTypeTable.C(activitytype.FieldTypeCode), "type"),
		sql.As(taskActivityTable.C(taskactivity.FieldUpdatedAt), "updated_at"),
	).From(taskActivityTable).LeftJoin(activityTypeTable).On(
		taskActivityTable.C(taskactivity.FieldActivityTypeID),
		activityTypeTable.C(activitytype.FieldID),
	).Where(sql.EQ(taskActivityTable.C(taskactivity.FieldWorkspaceID), where.WorkspaceID))

	workspaceActivitySelector := sql.Select(
		sql.As(workspaceActivityTable.C(workspaceactivity.FieldID), "id"),
		sql.As(activityTypeTable.C(activitytype.FieldTypeCode), "type"),
		sql.As(workspaceActivityTable.C(workspaceactivity.FieldUpdatedAt), "updated_at"),
	).From(workspaceActivityTable).LeftJoin(activityTypeTable).On(
		workspaceActivityTable.C(workspaceactivity.FieldActivityTypeID),
		activityTypeTable.C(activitytype.FieldID),
	).Where(sql.EQ(workspaceActivityTable.C(workspaceactivity.FieldWorkspaceID), where.WorkspaceID))

	t1 := sql.Table("t1").As("t1")
	selector := sql.Select(
		t1.C("id"),
		t1.C("type"),
		t1.C("updated_at"),
	).From(
		taskActivitySelector.Union(workspaceActivitySelector).As("t1"),
	).OrderBy(t1.C("updated_at"))

	queryString, args := selector.Query()
	rows, err := db.Query(queryString, args...)
	if err != nil {
		return nil, model.NewDBError(err)
	}

	defer rows.Close()

	var activities []*model.Activity
	for rows.Next() {
		var m model.Activity
		err = rows.Scan(&m.ID, &m.Type, &m.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		activities = append(activities, &m)
	}
	if rows.Err() != nil {
		return nil, model.NewDBError(err)
	}

	return activities, nil
}

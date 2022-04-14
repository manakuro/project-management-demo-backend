package seed

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/activitytype"
)

var activityTypeFeed = struct {
	task      ent.CreateActivityTypeInput
	workspace ent.CreateActivityTypeInput
}{
	task:      ent.CreateActivityTypeInput{Name: "Task", TypeCode: activitytype.TypeCodeTask},
	workspace: ent.CreateActivityTypeInput{Name: "Workspace", TypeCode: activitytype.TypeCodeWorkspace},
}

// ActivityType generates activity data.
func ActivityType(ctx context.Context, client *ent.Client) {
	_, err := client.ActivityType.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("ActivityType failed to delete data: %v", err)
	}

	data := []ent.CreateActivityTypeInput{
		activityTypeFeed.task,
		activityTypeFeed.workspace,
	}
	bulk := make([]*ent.ActivityTypeCreate, len(data))
	for i, t := range data {
		bulk[i] = client.ActivityType.Create().SetInput(t)
	}
	if _, err = client.ActivityType.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("ActivityType failed to seed data: %v", err)
	}
}

package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

var taskActivitySeed = make([]*ent.TaskActivity, 4)

// TaskActivity generates task activity data.
func TaskActivity(ctx context.Context, client *ent.Client) {
	_, err := client.TaskActivity.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TaskActivity failed to delete data: %v", err)
	}
	activityTypeID := seedutil.GetActivityType(ctx, client, "Task").ID
	teammateID := seedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID

	inputs := []ent.CreateTaskActivityInput{
		{
			TeammateID: teammateID,
			ActivityID: activityTypeID,
		},
		{
			TeammateID: teammateID,
			ActivityID: activityTypeID,
		},
		{
			TeammateID: teammateID,
			ActivityID: activityTypeID,
		},
		{
			TeammateID: teammateID,
			ActivityID: activityTypeID,
		},
	}
	bulk := make([]*ent.TaskActivityCreate, len(inputs))
	for i, t := range inputs {
		bulk[i] = client.TaskActivity.Create().SetInput(t)
	}
	if _, err = client.TaskActivity.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TaskActivity failed to feed data: %v", err)
	}

	res, err := client.TaskActivity.Query().All(ctx)
	if err != nil {
		log.Fatalf("TaskActivity failed to feed data: %v", err)
	}

	for i, r := range res {
		taskActivitySeed[i] = r
	}
}

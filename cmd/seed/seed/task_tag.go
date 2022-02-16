package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

// TaskTag generates tasks data
func TaskTag(ctx context.Context, client *ent.Client) {
	_, err := client.TaskTag.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TaskTag failed to delete data: %v", err)
	}
	designTag := seedutil.GetTagByName(ctx, client, tagFeed.design.Name)
	bugTag := seedutil.GetTagByName(ctx, client, tagFeed.bug.Name)
	developmentTag := seedutil.GetTagByName(ctx, client, tagFeed.development.Name)

	ts := []ent.CreateTaskTagInput{
		{
			TaskID: seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
			TagID:  developmentTag.ID,
		},
		{
			TaskID: seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2.Name).ID,
			TagID:  bugTag.ID,
		},
		{
			TaskID: seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task3.Name).ID,
			TagID:  designTag.ID,
		},
		{
			TaskID: seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task4.Name).ID,
			TagID:  developmentTag.ID,
		},
		{
			TaskID: seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task5.Name).ID,
			TagID:  developmentTag.ID,
		},
	}
	bulk := make([]*ent.TaskTagCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TaskTag.Create().SetInput(t)
	}
	if _, err = client.TaskTag.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TaskTag failed to feed data: %v", err)
	}
}

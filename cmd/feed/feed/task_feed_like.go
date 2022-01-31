package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// TaskFeedLike generates tasks data
func TaskFeedLike(ctx context.Context, client *ent.Client) {
	_, err := client.TaskFeedLike.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TaskFeedLike failed to delete data: %v", err)
	}
	manatoID := feedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID
	danID := feedutil.GetTeammateByEmail(ctx, client, teammateFeed.dan.Email).ID
	kentID := feedutil.GetTeammateByEmail(ctx, client, teammateFeed.kent.Email).ID

	task1ID := feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID
	taskFeeds := feedutil.GetTaskFeeds(ctx, client, task1ID)

	ts := []ent.CreateTaskFeedLikeInput{
		{
			TeammateID: manatoID,
			TaskID:     task1ID,
			TaskFeedID: taskFeeds[0].ID,
		},
		{
			TeammateID: danID,
			TaskID:     task1ID,
			TaskFeedID: taskFeeds[0].ID,
		},
		{
			TeammateID: kentID,
			TaskID:     task1ID,
			TaskFeedID: taskFeeds[0].ID,
		},
	}
	bulk := make([]*ent.TaskFeedLikeCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TaskFeedLike.Create().SetInput(t)
	}
	if _, err = client.TaskFeedLike.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TaskFeedLike failed to feed data: %v", err)
	}
}

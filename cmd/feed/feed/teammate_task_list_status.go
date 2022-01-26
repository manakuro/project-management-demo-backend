package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// TeammateTaskListStatus generates project_teammates data
func TeammateTaskListStatus(ctx context.Context, client *ent.Client) {
	_, err := client.TeammateTaskListStatus.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TeammateTaskListStatus failed to delete data: %v", err)
	}

	ts := []ent.CreateTeammateTaskListStatusInput{
		{
			WorkspaceID:               feedutil.GetWorkspace(ctx, client).ID,
			TeammateID:                feedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID,
			TaskListCompletedStatusID: feedutil.GetTaskListCompletedStatusByName(ctx, client, taskListCompletedStatusFeed.incomplete.Name).ID,
			TaskListSortStatusID:      feedutil.GetTaskListSortStatusByName(ctx, client, taskListSortStatusFeed.none.Name).ID,
		},
	}
	bulk := make([]*ent.TeammateTaskListStatusCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TeammateTaskListStatus.Create().SetInput(t)
	}
	if _, err = client.TeammateTaskListStatus.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TeammateTaskListStatus failed to feed data: %v", err)
	}
}

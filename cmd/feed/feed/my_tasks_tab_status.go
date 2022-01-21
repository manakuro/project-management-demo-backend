package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// MyTasksTabStatus generates project_teammates data
func MyTasksTabStatus(ctx context.Context, client *ent.Client) {
	_, err := client.MyTasksTabStatus.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("MyTasksTabStatus failed to delete data: %v", err)
	}

	ts := []ent.CreateMyTasksTabStatusInput{
		{
			WorkspaceID: feedutil.GetWorkspace(ctx, client).ID,
			TeammateID:  feedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID,
		},
	}
	bulk := make([]*ent.MyTasksTabStatusCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.MyTasksTabStatus.Create().SetInput(t)
	}
	if _, err = client.MyTasksTabStatus.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("MyTasksTabStatus failed to feed data: %v", err)
	}
}

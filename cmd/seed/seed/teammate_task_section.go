package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

var teammateTaskSectionFeed = struct {
	recentlyAssigned ent.CreateTeammateTaskSectionInput
	today            ent.CreateTeammateTaskSectionInput
}{
	recentlyAssigned: ent.CreateTeammateTaskSectionInput{
		Name:     "Recently assigned",
		Assigned: true,
	},
	today: ent.CreateTeammateTaskSectionInput{
		Name:     "Today",
		Assigned: false,
	},
}

// TeammateTaskSection generates project_teammates data
func TeammateTaskSection(ctx context.Context, client *ent.Client) {
	_, err := client.TeammateTaskSection.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TeammateTaskSection failed to delete data: %v", err)
	}

	ts := []ent.CreateTeammateTaskSectionInput{
		{
			WorkspaceID: seedutil.GetWorkspace(ctx, client).ID,
			TeammateID:  seedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID,
			Name:        teammateTaskSectionFeed.recentlyAssigned.Name,
			Assigned:    teammateTaskSectionFeed.recentlyAssigned.Assigned,
		},
		{
			WorkspaceID: seedutil.GetWorkspace(ctx, client).ID,
			TeammateID:  seedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID,
			Name:        teammateTaskSectionFeed.today.Name,
			Assigned:    teammateTaskSectionFeed.today.Assigned,
		},
		{
			WorkspaceID: seedutil.GetWorkspace(ctx, client).ID,
			TeammateID:  seedutil.GetTeammateByEmail(ctx, client, teammateFeed.dan.Email).ID,
			Name:        teammateTaskSectionFeed.recentlyAssigned.Name,
			Assigned:    teammateTaskSectionFeed.recentlyAssigned.Assigned,
		},
		{
			WorkspaceID: seedutil.GetWorkspace(ctx, client).ID,
			TeammateID:  seedutil.GetTeammateByEmail(ctx, client, teammateFeed.dan.Email).ID,
			Name:        teammateTaskSectionFeed.today.Name,
			Assigned:    teammateTaskSectionFeed.today.Assigned,
		},
		{
			WorkspaceID: seedutil.GetWorkspace(ctx, client).ID,
			TeammateID:  seedutil.GetTeammateByEmail(ctx, client, teammateFeed.kent.Email).ID,
			Name:        teammateTaskSectionFeed.recentlyAssigned.Name,
			Assigned:    teammateTaskSectionFeed.recentlyAssigned.Assigned,
		},
		{
			WorkspaceID: seedutil.GetWorkspace(ctx, client).ID,
			TeammateID:  seedutil.GetTeammateByEmail(ctx, client, teammateFeed.kent.Email).ID,
			Name:        teammateTaskSectionFeed.today.Name,
			Assigned:    teammateTaskSectionFeed.today.Assigned,
		},
	}
	bulk := make([]*ent.TeammateTaskSectionCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TeammateTaskSection.Create().SetInput(t)
	}
	if _, err = client.TeammateTaskSection.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TeammateTaskSection failed to feed data: %v", err)
	}
}

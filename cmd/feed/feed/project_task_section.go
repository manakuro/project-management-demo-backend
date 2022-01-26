package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

var projectTaskSectionFeed = struct {
	backlog    ent.CreateProjectTaskSectionInput
	ready      ent.CreateProjectTaskSectionInput
	inProgress ent.CreateProjectTaskSectionInput
	done       ent.CreateProjectTaskSectionInput
}{
	backlog: ent.CreateProjectTaskSectionInput{
		Name: "Backlog",
	},
	ready: ent.CreateProjectTaskSectionInput{
		Name: "Ready",
	},
	inProgress: ent.CreateProjectTaskSectionInput{
		Name: "In Progress",
	},
	done: ent.CreateProjectTaskSectionInput{
		Name: "Done",
	},
}

// ProjectTaskSection generates project_projects data
func ProjectTaskSection(ctx context.Context, client *ent.Client) {
	_, err := client.ProjectTaskSection.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("ProjectTaskSection failed to delete data: %v", err)
	}

	ts := []ent.CreateProjectTaskSectionInput{
		{
			ProjectID: feedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			Name:      projectTaskSectionFeed.backlog.Name,
		},
		{
			ProjectID: feedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			Name:      projectTaskSectionFeed.ready.Name,
		},
		{
			ProjectID: feedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			Name:      projectTaskSectionFeed.inProgress.Name,
		},
		{
			ProjectID: feedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			Name:      projectTaskSectionFeed.done.Name,
		},

		{
			ProjectID: feedutil.GetProjectByName(ctx, client, projectFeed.marketing.name).ID,
			Name:      projectTaskSectionFeed.backlog.Name,
		},
		{
			ProjectID: feedutil.GetProjectByName(ctx, client, projectFeed.marketing.name).ID,
			Name:      projectTaskSectionFeed.ready.Name,
		},
		{
			ProjectID: feedutil.GetProjectByName(ctx, client, projectFeed.marketing.name).ID,
			Name:      projectTaskSectionFeed.inProgress.Name,
		},
		{
			ProjectID: feedutil.GetProjectByName(ctx, client, projectFeed.marketing.name).ID,
			Name:      projectTaskSectionFeed.done.Name,
		},

		{
			ProjectID: feedutil.GetProjectByName(ctx, client, projectFeed.customerSuccess.name).ID,
			Name:      projectTaskSectionFeed.backlog.Name,
		},
		{
			ProjectID: feedutil.GetProjectByName(ctx, client, projectFeed.customerSuccess.name).ID,
			Name:      projectTaskSectionFeed.ready.Name,
		},
		{
			ProjectID: feedutil.GetProjectByName(ctx, client, projectFeed.customerSuccess.name).ID,
			Name:      projectTaskSectionFeed.inProgress.Name,
		},
		{
			ProjectID: feedutil.GetProjectByName(ctx, client, projectFeed.customerSuccess.name).ID,
			Name:      projectTaskSectionFeed.done.Name,
		},
	}
	bulk := make([]*ent.ProjectTaskSectionCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.ProjectTaskSection.Create().SetInput(t)
	}
	if _, err = client.ProjectTaskSection.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("ProjectTaskSection failed to feed data: %v", err)
	}
}

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

	appDevelopmentProject := feedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name)
	marketingProject := feedutil.GetProjectByName(ctx, client, projectFeed.marketing.name)
	customerSuccessProject := feedutil.GetProjectByName(ctx, client, projectFeed.customerSuccess.name)

	ts := []ent.CreateProjectTaskSectionInput{
		{
			ProjectID: appDevelopmentProject.ID,
			Name:      projectTaskSectionFeed.backlog.Name,
		},
		{
			ProjectID: appDevelopmentProject.ID,
			Name:      projectTaskSectionFeed.ready.Name,
		},
		{
			ProjectID: appDevelopmentProject.ID,
			Name:      projectTaskSectionFeed.inProgress.Name,
		},
		{
			ProjectID: appDevelopmentProject.ID,
			Name:      projectTaskSectionFeed.done.Name,
		},

		{
			ProjectID: marketingProject.ID,
			Name:      projectTaskSectionFeed.backlog.Name,
		},
		{
			ProjectID: marketingProject.ID,
			Name:      projectTaskSectionFeed.ready.Name,
		},
		{
			ProjectID: marketingProject.ID,
			Name:      projectTaskSectionFeed.inProgress.Name,
		},
		{
			ProjectID: marketingProject.ID,
			Name:      projectTaskSectionFeed.done.Name,
		},

		{
			ProjectID: customerSuccessProject.ID,
			Name:      projectTaskSectionFeed.backlog.Name,
		},
		{
			ProjectID: customerSuccessProject.ID,
			Name:      projectTaskSectionFeed.ready.Name,
		},
		{
			ProjectID: customerSuccessProject.ID,
			Name:      projectTaskSectionFeed.inProgress.Name,
		},
		{
			ProjectID: customerSuccessProject.ID,
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

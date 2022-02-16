package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

var projectTaskSectionFeedAppDevelopment = struct {
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

var projectTaskSectionFeedMarketing = struct {
	upcoming           ent.CreateProjectTaskSectionInput
	planning           ent.CreateProjectTaskSectionInput
	campaign           ent.CreateProjectTaskSectionInput
	contentDevelopment ent.CreateProjectTaskSectionInput
}{
	upcoming: ent.CreateProjectTaskSectionInput{
		Name: "Upcoming Work",
	},
	planning: ent.CreateProjectTaskSectionInput{
		Name: "Planning",
	},
	contentDevelopment: ent.CreateProjectTaskSectionInput{
		Name: "Content Development",
	},
	campaign: ent.CreateProjectTaskSectionInput{
		Name: "Campaign Promotion",
	},
}

// ProjectTaskSection generates project_projects data
func ProjectTaskSection(ctx context.Context, client *ent.Client) {
	_, err := client.ProjectTaskSection.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("ProjectTaskSection failed to delete data: %v", err)
	}

	appDevelopmentProject := seedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name)
	marketingProject := seedutil.GetProjectByName(ctx, client, projectFeed.marketing.name)
	customerSuccessProject := seedutil.GetProjectByName(ctx, client, projectFeed.customerSuccess.name)

	ts := []ent.CreateProjectTaskSectionInput{
		// App Development
		{
			ProjectID: appDevelopmentProject.ID,
			Name:      projectTaskSectionFeedAppDevelopment.backlog.Name,
		},
		{
			ProjectID: appDevelopmentProject.ID,
			Name:      projectTaskSectionFeedAppDevelopment.ready.Name,
		},
		{
			ProjectID: appDevelopmentProject.ID,
			Name:      projectTaskSectionFeedAppDevelopment.inProgress.Name,
		},
		{
			ProjectID: appDevelopmentProject.ID,
			Name:      projectTaskSectionFeedAppDevelopment.done.Name,
		},

		// Marketing
		{
			ProjectID: marketingProject.ID,
			Name:      projectTaskSectionFeedMarketing.upcoming.Name,
		},
		{
			ProjectID: marketingProject.ID,
			Name:      projectTaskSectionFeedMarketing.planning.Name,
		},
		{
			ProjectID: marketingProject.ID,
			Name:      projectTaskSectionFeedMarketing.contentDevelopment.Name,
		},
		{
			ProjectID: marketingProject.ID,
			Name:      projectTaskSectionFeedMarketing.campaign.Name,
		},

		// Customer Success
		{
			ProjectID: customerSuccessProject.ID,
			Name:      projectTaskSectionFeedAppDevelopment.backlog.Name,
		},
		{
			ProjectID: customerSuccessProject.ID,
			Name:      projectTaskSectionFeedAppDevelopment.ready.Name,
		},
		{
			ProjectID: customerSuccessProject.ID,
			Name:      projectTaskSectionFeedAppDevelopment.inProgress.Name,
		},
		{
			ProjectID: customerSuccessProject.ID,
			Name:      projectTaskSectionFeedAppDevelopment.done.Name,
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

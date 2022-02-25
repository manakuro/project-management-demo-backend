package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

// ProjectTaskColumn generates project_teammates data
func ProjectTaskColumn(ctx context.Context, client *ent.Client) {
	_, err := client.ProjectTaskColumn.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("ProjectTaskColumn failed to delete data: %v", err)
	}

	ts := []ent.CreateProjectTaskColumnInput{
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.taskName.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			Width:        "500px",
			Disabled:     false,
			Customizable: false,
			Order:        1,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.assignee.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        2,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.dueDate.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        3,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.tags.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        4,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.priority.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        5,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.taskName.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.marketing.name).ID,
			Width:        "500px",
			Disabled:     false,
			Customizable: false,
			Order:        1,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.assignee.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.marketing.name).ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        2,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.dueDate.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.marketing.name).ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        3,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.tags.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.marketing.name).ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        4,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.priority.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.marketing.name).ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        5,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.taskName.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.customerSuccess.name).ID,
			Width:        "500px",
			Disabled:     false,
			Customizable: false,
			Order:        1,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.assignee.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.customerSuccess.name).ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        2,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.dueDate.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.customerSuccess.name).ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        3,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.tags.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.customerSuccess.name).ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        4,
		},
		{
			TaskColumnID: seedutil.GetTaskColumnByName(ctx, client, taskColumnFeed.priority.Name).ID,
			ProjectID:    seedutil.GetProjectByName(ctx, client, projectFeed.customerSuccess.name).ID,
			Width:        "120px",
			Disabled:     false,
			Customizable: true,
			Order:        5,
		},
	}
	bulk := make([]*ent.ProjectTaskColumnCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.ProjectTaskColumn.Create().SetInput(t)
	}
	if _, err = client.ProjectTaskColumn.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("ProjectTaskColumn failed to feed data: %v", err)
	}
}

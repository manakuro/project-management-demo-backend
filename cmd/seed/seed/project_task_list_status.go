package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

// ProjectTaskListStatus generates project_projects data
func ProjectTaskListStatus(ctx context.Context, client *ent.Client) {
	_, err := client.ProjectTaskListStatus.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("ProjectTaskListStatus failed to delete data: %v", err)
	}

	ts := []ent.CreateProjectTaskListStatusInput{
		{
			ProjectID:                 seedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID,
			TaskListCompletedStatusID: seedutil.GetTaskListCompletedStatusByName(ctx, client, taskListCompletedStatusFeed.incomplete.Name).ID,
			TaskListSortStatusID:      seedutil.GetTaskListSortStatusByName(ctx, client, taskListSortStatusFeed.none.Name).ID,
		},
		{
			ProjectID:                 seedutil.GetProjectByName(ctx, client, projectFeed.marketing.name).ID,
			TaskListCompletedStatusID: seedutil.GetTaskListCompletedStatusByName(ctx, client, taskListCompletedStatusFeed.incomplete.Name).ID,
			TaskListSortStatusID:      seedutil.GetTaskListSortStatusByName(ctx, client, taskListSortStatusFeed.none.Name).ID,
		},
		{
			ProjectID:                 seedutil.GetProjectByName(ctx, client, projectFeed.customerSuccess.name).ID,
			TaskListCompletedStatusID: seedutil.GetTaskListCompletedStatusByName(ctx, client, taskListCompletedStatusFeed.incomplete.Name).ID,
			TaskListSortStatusID:      seedutil.GetTaskListSortStatusByName(ctx, client, taskListSortStatusFeed.none.Name).ID,
		},
	}
	bulk := make([]*ent.ProjectTaskListStatusCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.ProjectTaskListStatus.Create().SetInput(t)
	}
	if _, err = client.ProjectTaskListStatus.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("ProjectTaskListStatus failed to feed data: %v", err)
	}
}

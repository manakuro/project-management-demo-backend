package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// ProjectTask generates tasks data
func ProjectTask(ctx context.Context, client *ent.Client) {
	_, err := client.ProjectTask.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("ProjectTask failed to delete data: %v", err)
	}

	appDevelopmentID := feedutil.GetProjectByName(ctx, client, projectFeed.appDevelopment.name).ID
	ts := []ent.CreateProjectTaskInput{
		// appDevelopment
		{
			ProjectID:            appDevelopmentID,
			TaskID:               feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
			ProjectTaskSectionID: feedutil.GetProjectTaskSectionByName(ctx, client, appDevelopmentID, projectTaskSectionFeed.backlog.Name).ID,
		},
		{
			ProjectID:            appDevelopmentID,
			TaskID:               feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2.Name).ID,
			ProjectTaskSectionID: feedutil.GetProjectTaskSectionByName(ctx, client, appDevelopmentID, projectTaskSectionFeed.backlog.Name).ID,
		},
		{
			ProjectID:            appDevelopmentID,
			TaskID:               feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task3.Name).ID,
			ProjectTaskSectionID: feedutil.GetProjectTaskSectionByName(ctx, client, appDevelopmentID, projectTaskSectionFeed.backlog.Name).ID,
		},
		{
			ProjectID:            appDevelopmentID,
			TaskID:               feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task4.Name).ID,
			ProjectTaskSectionID: feedutil.GetProjectTaskSectionByName(ctx, client, appDevelopmentID, projectTaskSectionFeed.ready.Name).ID,
		},
		{
			ProjectID:            appDevelopmentID,
			TaskID:               feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task5.Name).ID,
			ProjectTaskSectionID: feedutil.GetProjectTaskSectionByName(ctx, client, appDevelopmentID, projectTaskSectionFeed.ready.Name).ID,
		},
		{
			ProjectID:            appDevelopmentID,
			TaskID:               feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task6.Name).ID,
			ProjectTaskSectionID: feedutil.GetProjectTaskSectionByName(ctx, client, appDevelopmentID, projectTaskSectionFeed.ready.Name).ID,
		},
		{
			ProjectID:            appDevelopmentID,
			TaskID:               feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task7.Name).ID,
			ProjectTaskSectionID: feedutil.GetProjectTaskSectionByName(ctx, client, appDevelopmentID, projectTaskSectionFeed.inProgress.Name).ID,
		},
		{
			ProjectID:            appDevelopmentID,
			TaskID:               feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task8.Name).ID,
			ProjectTaskSectionID: feedutil.GetProjectTaskSectionByName(ctx, client, appDevelopmentID, projectTaskSectionFeed.inProgress.Name).ID,
		},
		{
			ProjectID:            appDevelopmentID,
			TaskID:               feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task9.Name).ID,
			ProjectTaskSectionID: feedutil.GetProjectTaskSectionByName(ctx, client, appDevelopmentID, projectTaskSectionFeed.inProgress.Name).ID,
		},
		{
			ProjectID:            appDevelopmentID,
			TaskID:               feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task10.Name).ID,
			ProjectTaskSectionID: feedutil.GetProjectTaskSectionByName(ctx, client, appDevelopmentID, projectTaskSectionFeed.inProgress.Name).ID,
		},
		{
			ProjectID:            appDevelopmentID,
			TaskID:               feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2Subtask1.Name).ID,
			ProjectTaskSectionID: feedutil.GetProjectTaskSectionByName(ctx, client, appDevelopmentID, projectTaskSectionFeed.inProgress.Name).ID,
		},
		{
			ProjectID:            appDevelopmentID,
			TaskID:               feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2Subtask2.Name).ID,
			ProjectTaskSectionID: feedutil.GetProjectTaskSectionByName(ctx, client, appDevelopmentID, projectTaskSectionFeed.done.Name).ID,
		},
		{
			ProjectID:            appDevelopmentID,
			TaskID:               feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2Subtask3.Name).ID,
			ProjectTaskSectionID: feedutil.GetProjectTaskSectionByName(ctx, client, appDevelopmentID, projectTaskSectionFeed.done.Name).ID,
		},
	}
	bulk := make([]*ent.ProjectTaskCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.ProjectTask.Create().SetInput(t)
	}
	if _, err = client.ProjectTask.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("ProjectTask failed to feed data: %v", err)
	}
}

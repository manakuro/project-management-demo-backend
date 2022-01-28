package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// TeammateTask generates tasks data
func TeammateTask(ctx context.Context, client *ent.Client) {
	_, err := client.TeammateTask.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TeammateTask failed to delete data: %v", err)
	}
	teammate := feedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email)

	ts := []ent.CreateTeammateTaskInput{
		{
			TeammateID:            teammate.ID,
			TaskID:                feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
			TeammateTaskSectionID: feedutil.GetTeammateTaskSectionByName(ctx, client, teammateTaskSectionFeed.recentlyAssigned.Name).ID,
		},
		{
			TeammateID:            teammate.ID,
			TaskID:                feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2.Name).ID,
			TeammateTaskSectionID: feedutil.GetTeammateTaskSectionByName(ctx, client, teammateTaskSectionFeed.recentlyAssigned.Name).ID,
		},
		{
			TeammateID:            teammate.ID,
			TaskID:                feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task3.Name).ID,
			TeammateTaskSectionID: feedutil.GetTeammateTaskSectionByName(ctx, client, teammateTaskSectionFeed.recentlyAssigned.Name).ID,
		},
		{
			TeammateID:            teammate.ID,
			TaskID:                feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task4.Name).ID,
			TeammateTaskSectionID: feedutil.GetTeammateTaskSectionByName(ctx, client, teammateTaskSectionFeed.recentlyAssigned.Name).ID,
		},
		{
			TeammateID:            teammate.ID,
			TaskID:                feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task5.Name).ID,
			TeammateTaskSectionID: feedutil.GetTeammateTaskSectionByName(ctx, client, teammateTaskSectionFeed.recentlyAssigned.Name).ID,
		},
		{
			TeammateID:            teammate.ID,
			TaskID:                feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task6.Name).ID,
			TeammateTaskSectionID: feedutil.GetTeammateTaskSectionByName(ctx, client, teammateTaskSectionFeed.recentlyAssigned.Name).ID,
		},
		{
			TeammateID:            teammate.ID,
			TaskID:                feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task7.Name).ID,
			TeammateTaskSectionID: feedutil.GetTeammateTaskSectionByName(ctx, client, teammateTaskSectionFeed.recentlyAssigned.Name).ID,
		},
		{
			TeammateID:            teammate.ID,
			TaskID:                feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task8.Name).ID,
			TeammateTaskSectionID: feedutil.GetTeammateTaskSectionByName(ctx, client, teammateTaskSectionFeed.recentlyAssigned.Name).ID,
		},
		{
			TeammateID:            teammate.ID,
			TaskID:                feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task9.Name).ID,
			TeammateTaskSectionID: feedutil.GetTeammateTaskSectionByName(ctx, client, teammateTaskSectionFeed.recentlyAssigned.Name).ID,
		},
		{
			TeammateID:            teammate.ID,
			TaskID:                feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task10.Name).ID,
			TeammateTaskSectionID: feedutil.GetTeammateTaskSectionByName(ctx, client, teammateTaskSectionFeed.today.Name).ID,
		},
		{
			TeammateID:            teammate.ID,
			TaskID:                feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2Subtask1.Name).ID,
			TeammateTaskSectionID: feedutil.GetTeammateTaskSectionByName(ctx, client, teammateTaskSectionFeed.today.Name).ID,
		},
		{
			TeammateID:            teammate.ID,
			TaskID:                feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2Subtask2.Name).ID,
			TeammateTaskSectionID: feedutil.GetTeammateTaskSectionByName(ctx, client, teammateTaskSectionFeed.today.Name).ID,
		},
		{
			TeammateID:            teammate.ID,
			TaskID:                feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2Subtask3.Name).ID,
			TeammateTaskSectionID: feedutil.GetTeammateTaskSectionByName(ctx, client, teammateTaskSectionFeed.today.Name).ID,
		},
	}
	bulk := make([]*ent.TeammateTaskCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TeammateTask.Create().SetInput(t)
	}
	if _, err = client.TeammateTask.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TeammateTask failed to feed data: %v", err)
	}
}

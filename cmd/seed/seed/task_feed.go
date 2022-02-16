package seed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/seed/seedutil"
	"project-management-demo-backend/ent"
)

// TaskFeed generates tasks data
func TaskFeed(ctx context.Context, client *ent.Client) {
	_, err := client.TaskFeed.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TaskFeed failed to delete data: %v", err)
	}
	manatoID := seedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID

	isFirstTrue := true

	ts := []ent.CreateTaskFeedInput{
		// task 1
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
			TeammateID:  manatoID,
			Description: seedutil.ParseDescription([]byte(`{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Take a look at the attachment above.","attrs":{"mentionId":"","mentionType":""}}]}]}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
			TeammateID:  manatoID,
			Description: seedutil.ParseDescription([]byte(`{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Could you give me a review for the proposal?","attrs":{"mentionId":"","mentionType":""}}]}]}`)),
		},

		// task 2
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 3
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task3.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 4
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task4.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 5
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task5.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 6
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task6.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 7
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task7.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 8
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task8.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 9
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task9.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 10
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task10.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 2 - 1
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2Subtask1.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 2 - 2
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2Subtask2.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 2 - 3
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskAssignedFeed.task2Subtask3.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},

		/** taskNoAssignedFeed **/
		// task 1
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.task1.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 2
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.task2.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 2 - 1
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.task2Subtask1.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 2 - 2
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.task2Subtask2.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		// task 2 - 3
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.task2Subtask3.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},

		/** Marketing **/
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask1.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask2.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask3.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask4.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask5.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask6.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask7.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask8.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask9.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask10.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask11.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask12.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask13.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      seedutil.GetTaskByName(ctx, client, taskNoAssignedFeed.mTask14.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: seedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
	}
	bulk := make([]*ent.TaskFeedCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TaskFeed.Create().SetInput(t)
	}
	if _, err = client.TaskFeed.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TaskFeed failed to feed data: %v", err)
	}
}

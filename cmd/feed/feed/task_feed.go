package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// TaskFeed generates tasks data
func TaskFeed(ctx context.Context, client *ent.Client) {
	_, err := client.TaskFeed.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TaskFeed failed to delete data: %v", err)
	}
	manatoID := feedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email).ID

	isFirstTrue := true

	ts := []ent.CreateTaskFeedInput{
		{
			TaskID:      feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
			TeammateID:  manatoID,
			IsFirst:     &isFirstTrue,
			Description: feedutil.ParseDescription([]byte(`{"type": "doc", "content": []}`)),
		},
		{
			TaskID:      feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
			TeammateID:  manatoID,
			Description: feedutil.ParseDescription([]byte(`{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Welcome to the Marketing team! We’ll be using this project to track our progress on our Q1 product launch. Final ad designs are in the “Key Resources” section below. Use this form to submit new ideas!"}]},{"type":"paragraph"},{"type":"paragraph","content":[{"type":"text","text":"Project Owner: "},{"type":"mention","attrs":{"mentionId":"1","mentionType":"1"}}]},{"type":"paragraph","content":[{"type":"text","text":"Tech Lead:"}]},{"type":"paragraph","content":[{"type":"mention","attrs":{"mentionId":"2","mentionType":"1"}}]},{"type":"paragraph"},{"type":"paragraph"}]}`)),
		},
		{
			TaskID:      feedutil.GetTaskByName(ctx, client, taskAssignedFeed.task1.Name).ID,
			TeammateID:  manatoID,
			Description: feedutil.ParseDescription([]byte(`{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"😜"}]},{"type":"paragraph","content":[{"type":"text","text":"test"}]}]}`)),
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

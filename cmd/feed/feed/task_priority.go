package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/taskpriority"
)

var taskPriorityFeed = struct {
	low    ent.CreateTaskPriorityInput
	medium ent.CreateTaskPriorityInput
	high   ent.CreateTaskPriorityInput
}{
	low: ent.CreateTaskPriorityInput{
		Name:         "Low",
		PriorityType: taskpriority.PriorityTypeLow,
	},
	medium: ent.CreateTaskPriorityInput{
		Name:         "Medium",
		PriorityType: taskpriority.PriorityTypeMedium,
	},
	high: ent.CreateTaskPriorityInput{
		Name:         "High",
		PriorityType: taskpriority.PriorityTypeHigh,
	},
}

// TaskPriority generates project light color data
func TaskPriority(ctx context.Context, client *ent.Client) {
	_, err := client.TaskPriority.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TaskPriority failed to delete data: %v", err)
	}

	data := []ent.CreateTaskPriorityInput{
		{
			Name:         taskPriorityFeed.low.Name,
			PriorityType: taskPriorityFeed.low.PriorityType,
			ColorID:      feedutil.GetColor(ctx, client, colorFeed.green400.Color).ID,
		},
		{
			Name:         taskPriorityFeed.medium.Name,
			PriorityType: taskPriorityFeed.medium.PriorityType,
			ColorID:      feedutil.GetColor(ctx, client, colorFeed.orange400.Color).ID,
		},
		{
			Name:         taskPriorityFeed.high.Name,
			PriorityType: taskPriorityFeed.high.PriorityType,
			ColorID:      feedutil.GetColor(ctx, client, colorFeed.red400.Color).ID,
		},
	}
	bulk := make([]*ent.TaskPriorityCreate, len(data))
	for i, t := range data {
		bulk[i] = client.TaskPriority.Create().SetInput(t)
	}
	if _, err = client.TaskPriority.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TaskPriority failed to feed data: %v", err)
	}
}

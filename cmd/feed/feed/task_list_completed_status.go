package feed

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/tasklistcompletedstatus"
)

var taskListCompletedStatusFeed = struct {
	incomplete         ent.CreateTaskListCompletedStatusInput
	completed          ent.CreateTaskListCompletedStatusInput
	completedToday     ent.CreateTaskListCompletedStatusInput
	completedYesterday ent.CreateTaskListCompletedStatusInput
	completed1Week     ent.CreateTaskListCompletedStatusInput
	completed2Weeks    ent.CreateTaskListCompletedStatusInput
	completed3Weeks    ent.CreateTaskListCompletedStatusInput
	all                ent.CreateTaskListCompletedStatusInput
}{
	incomplete:         ent.CreateTaskListCompletedStatusInput{Name: "Incomplete", StatusCode: tasklistcompletedstatus.StatusCodeIncomplete},
	completed:          ent.CreateTaskListCompletedStatusInput{Name: "Completed", StatusCode: tasklistcompletedstatus.StatusCodeCompleted},
	completedToday:     ent.CreateTaskListCompletedStatusInput{Name: "Completed Today", StatusCode: tasklistcompletedstatus.StatusCodeCompletedToday},
	completedYesterday: ent.CreateTaskListCompletedStatusInput{Name: "Completed Yesterday", StatusCode: tasklistcompletedstatus.StatusCodeCompletedYesterday},
	completed1Week:     ent.CreateTaskListCompletedStatusInput{Name: "Completed 1 Week", StatusCode: tasklistcompletedstatus.StatusCodeCompleted1Week},
	completed2Weeks:    ent.CreateTaskListCompletedStatusInput{Name: "Completed 2 Weeks", StatusCode: tasklistcompletedstatus.StatusCodeCompleted2Weeks},
	completed3Weeks:    ent.CreateTaskListCompletedStatusInput{Name: "Completed 3 Weeks", StatusCode: tasklistcompletedstatus.StatusCodeCompleted3Weeks},
	all:                ent.CreateTaskListCompletedStatusInput{Name: "All", StatusCode: tasklistcompletedstatus.StatusCodeAll},
}

// TaskListCompletedStatus generates color data
func TaskListCompletedStatus(ctx context.Context, client *ent.Client) {
	_, err := client.TaskListCompletedStatus.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TaskListCompletedStatus failed to delete data: %v", err)
	}

	ts := []ent.CreateTaskListCompletedStatusInput{
		taskListCompletedStatusFeed.incomplete,
		taskListCompletedStatusFeed.completed,
		taskListCompletedStatusFeed.completedToday,
		taskListCompletedStatusFeed.completedYesterday,
		taskListCompletedStatusFeed.completed1Week,
		taskListCompletedStatusFeed.completed2Weeks,
		taskListCompletedStatusFeed.completed3Weeks,
		taskListCompletedStatusFeed.all,
	}
	bulk := make([]*ent.TaskListCompletedStatusCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TaskListCompletedStatus.Create().SetInput(t)
	}
	if _, err = client.TaskListCompletedStatus.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TaskListCompletedStatus failed to feed data: %v", err)
	}
}

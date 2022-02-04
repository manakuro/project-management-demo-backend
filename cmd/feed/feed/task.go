package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

var taskAssignedFeed = struct {
	task1         ent.CreateTaskInput
	task2         ent.CreateTaskInput
	task3         ent.CreateTaskInput
	task4         ent.CreateTaskInput
	task5         ent.CreateTaskInput
	task6         ent.CreateTaskInput
	task7         ent.CreateTaskInput
	task8         ent.CreateTaskInput
	task9         ent.CreateTaskInput
	task10        ent.CreateTaskInput
	task2Subtask1 ent.CreateTaskInput
	task2Subtask2 ent.CreateTaskInput
	task2Subtask3 ent.CreateTaskInput
}{
	task1:         ent.CreateTaskInput{Name: "Implement new card design", DueDate: feedutil.AddDate(10)},
	task2:         ent.CreateTaskInput{Name: "User bug report", DueDate: feedutil.AddDate(5)},
	task2Subtask1: ent.CreateTaskInput{Name: "User getting sent duplicate notifications", DueDate: feedutil.AddDate(2)},
	task2Subtask2: ent.CreateTaskInput{Name: "User can't invite teammate via modal page", DueDate: feedutil.AddDate(2)},
	task2Subtask3: ent.CreateTaskInput{Name: "Broken links on my page", DueDate: feedutil.AddDate(2)},
	task3:         ent.CreateTaskInput{Name: "Design iOS prototype"},
	task4:         ent.CreateTaskInput{Name: "Scope performance improvements"},
	task5:         ent.CreateTaskInput{Name: "Implement mobile menu"},
	task6:         ent.CreateTaskInput{Name: "Support for offline mode"},
	task7:         ent.CreateTaskInput{Name: "Introduce CI"},
	task8:         ent.CreateTaskInput{Name: "Login with Google"},
	task9:         ent.CreateTaskInput{Name: "Implement undo function"},
	task10:        ent.CreateTaskInput{Name: "Export to PDF file"},
}

var taskNoAssignedFeed = struct {
	task1         ent.CreateTaskInput
	task2         ent.CreateTaskInput
	task2Subtask1 ent.CreateTaskInput
	task2Subtask2 ent.CreateTaskInput
	task2Subtask3 ent.CreateTaskInput
}{
	task1:         ent.CreateTaskInput{Name: "Launch updated task list", DueDate: feedutil.AddDate(10)},
	task2:         ent.CreateTaskInput{Name: "Finalize workspace design", DueDate: feedutil.AddDate(5)},
	task2Subtask1: ent.CreateTaskInput{Name: "Prep for review", DueDate: feedutil.AddDate(2)},
	task2Subtask2: ent.CreateTaskInput{Name: "Tweak project card design", DueDate: feedutil.AddDate(2)},
	task2Subtask3: ent.CreateTaskInput{Name: "Usability testing for top bar", DueDate: feedutil.AddDate(2)},
}

// Task generates tasks data
func Task(ctx context.Context, client *ent.Client) {
	_, err := client.Task.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("Task failed to delete data: %v", err)
	}
	teammate := feedutil.GetTeammateByEmail(ctx, client, teammateFeed.manato.Email)
	completed := true

	assignedTask2, err := client.Task.Create().SetInput(ent.CreateTaskInput{
		Name:           taskAssignedFeed.task2.Name,
		DueDate:        taskAssignedFeed.task2.DueDate,
		AssigneeID:     &teammate.ID,
		CreatedBy:      teammate.ID,
		TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.high.Name).ID,
	}).Save(ctx)
	if err != nil {
		log.Fatalf("Task failed to feed data: %v", err)
	}

	noAssignedTask2, err := client.Task.Create().SetInput(ent.CreateTaskInput{
		Name:           taskNoAssignedFeed.task2.Name,
		DueDate:        taskNoAssignedFeed.task2.DueDate,
		CreatedBy:      teammate.ID,
		TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.high.Name).ID,
	}).Save(ctx)
	if err != nil {
		log.Fatalf("Task failed to feed data: %v", err)
	}

	ts := []ent.CreateTaskInput{
		{
			Name:           taskAssignedFeed.task1.Name,
			DueDate:        taskAssignedFeed.task1.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.medium.Name).ID,
		},

		{
			Name:           taskAssignedFeed.task2Subtask1.Name,
			DueDate:        taskAssignedFeed.task2Subtask1.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskParentID:   &assignedTask2.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.medium.Name).ID,
		},
		{
			Name:           taskAssignedFeed.task2Subtask2.Name,
			DueDate:        taskAssignedFeed.task2Subtask2.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskParentID:   &assignedTask2.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.medium.Name).ID,
		},
		{
			Name:           taskAssignedFeed.task2Subtask3.Name,
			DueDate:        taskAssignedFeed.task2Subtask3.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskParentID:   &assignedTask2.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.medium.Name).ID,
		},

		{
			Name:           taskAssignedFeed.task3.Name,
			DueDate:        taskAssignedFeed.task3.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.medium.Name).ID,
		},
		{
			Name:           taskAssignedFeed.task4.Name,
			DueDate:        taskAssignedFeed.task4.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.medium.Name).ID,
		},
		{
			Name:           taskAssignedFeed.task5.Name,
			DueDate:        taskAssignedFeed.task5.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.low.Name).ID,
		},
		{
			Name:           taskAssignedFeed.task6.Name,
			DueDate:        taskAssignedFeed.task6.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.low.Name).ID,
		},
		{
			Name:           taskAssignedFeed.task7.Name,
			DueDate:        taskAssignedFeed.task7.DueDate,
			AssigneeID:     &teammate.ID,
			CreatedBy:      teammate.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.low.Name).ID,
		},
		{
			Name:           taskAssignedFeed.task8.Name,
			DueDate:        taskAssignedFeed.task8.DueDate,
			AssigneeID:     &teammate.ID,
			Completed:      &completed,
			CompletedAt:    feedutil.AddDate(-2),
			CreatedBy:      teammate.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.low.Name).ID,
		},
		{
			Name:           taskAssignedFeed.task9.Name,
			DueDate:        taskAssignedFeed.task9.DueDate,
			AssigneeID:     &teammate.ID,
			Completed:      &completed,
			CompletedAt:    feedutil.AddDate(-10),
			CreatedBy:      teammate.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.low.Name).ID,
		},
		{
			Name:           taskAssignedFeed.task10.Name,
			DueDate:        taskAssignedFeed.task10.DueDate,
			AssigneeID:     &teammate.ID,
			Completed:      &completed,
			CompletedAt:    feedutil.AddDate(-7),
			CreatedBy:      teammate.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.low.Name).ID,
		},

		// No assigned feed
		{
			Name:           taskNoAssignedFeed.task2Subtask1.Name,
			DueDate:        taskAssignedFeed.task2Subtask1.DueDate,
			CreatedBy:      teammate.ID,
			TaskParentID:   &noAssignedTask2.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.medium.Name).ID,
		},
		{
			Name:           taskNoAssignedFeed.task2Subtask2.Name,
			DueDate:        taskAssignedFeed.task2Subtask2.DueDate,
			CreatedBy:      teammate.ID,
			TaskParentID:   &noAssignedTask2.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.medium.Name).ID,
		},
		{
			Name:           taskNoAssignedFeed.task2Subtask3.Name,
			DueDate:        taskAssignedFeed.task2Subtask3.DueDate,
			CreatedBy:      teammate.ID,
			TaskParentID:   &noAssignedTask2.ID,
			TaskPriorityID: feedutil.GetTaskPriorityByName(ctx, client, taskPriorityFeed.medium.Name).ID,
		},
	}
	bulk := make([]*ent.TaskCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.Task.Create().SetInput(t)
	}
	if _, err = client.Task.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("Task failed to feed data: %v", err)
	}
}

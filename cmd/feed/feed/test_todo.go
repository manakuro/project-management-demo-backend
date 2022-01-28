package feed

import (
	"context"
	"log"
	"project-management-demo-backend/cmd/feed/feedutil"
	"project-management-demo-backend/ent"
)

// TestTodo generates test todo data
func TestTodo(ctx context.Context, client *ent.Client) {
	_, err := client.TestTodo.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("TestTodo failed to delete data: %v", err)
	}
	id := feedutil.GetTestUserByName(ctx, client, testUserFeed.tom.Name).ID

	parentTodo, err := client.TestTodo.Create().SetInput(ent.CreateTestTodoInput{
		TestUserID: &id,
	}).Save(ctx)
	if err != nil {
		log.Fatalf("TestTodo failed to feed data: %v", err)
	}

	c1, err := client.TestTodo.
		Create().
		SetInput(ent.CreateTestTodoInput{
			TestUserID: &id,
			CreatedBy:  id,
		}).
		SetParent(parentTodo).
		Save(ctx)
	if err != nil {
		log.Fatalf("TestTodo failed to feed data: %v", err)
	}

	_, err = client.TestTodo.
		Create().
		SetTestUserID(id).
		SetParent(c1).
		SetCreatedBy(id).
		Save(ctx)
	if err != nil {
		log.Fatalf("TestTodo failed to feed data: %v", err)
	}

	ts := []ent.CreateTestTodoInput{
		{TestUserID: &id, ParentTodoID: &parentTodo.ID, CreatedBy: id},
		{TestUserID: &id, ParentTodoID: &parentTodo.ID, CreatedBy: id},
		{TestUserID: &id, ParentTodoID: &parentTodo.ID, CreatedBy: id},
		{TestUserID: &id, ParentTodoID: &parentTodo.ID, CreatedBy: id},
		{TestUserID: &id, ParentTodoID: &parentTodo.ID, CreatedBy: id},
		{TestUserID: &id, ParentTodoID: &parentTodo.ID, CreatedBy: id},
		{TestUserID: &id, ParentTodoID: &parentTodo.ID, CreatedBy: id},
		{TestUserID: &id, ParentTodoID: &parentTodo.ID, CreatedBy: id},
		{TestUserID: &id, ParentTodoID: &parentTodo.ID, CreatedBy: id},
		{TestUserID: &id, ParentTodoID: &parentTodo.ID, CreatedBy: id},
		{TestUserID: &id, ParentTodoID: &parentTodo.ID, CreatedBy: id},
		{TestUserID: &id, ParentTodoID: &parentTodo.ID, CreatedBy: id},
	}
	bulk := make([]*ent.TestTodoCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TestTodo.Create().SetInput(t)
	}
	if _, err = client.TestTodo.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("TestTodo failed to feed data: %v", err)
	}
}

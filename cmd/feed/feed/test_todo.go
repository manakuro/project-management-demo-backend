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
		log.Fatalf("test_todo: failed to delete test todo: %v", err)
	}

	id := feedutil.GetTestUserByName(ctx, client, testUserFeed.tom.Name).ID
	ts := []ent.CreateTestTodoInput{
		{TestUserID: &id},
		{TestUserID: &id},
		{TestUserID: &id},
		{TestUserID: &id},
		{TestUserID: &id},
		{TestUserID: &id},
		{TestUserID: &id},
		{TestUserID: &id},
		{TestUserID: &id},
		{TestUserID: &id},
		{TestUserID: &id},
		{TestUserID: &id},
	}
	bulk := make([]*ent.TestTodoCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TestTodo.Create().SetInput(t)
	}
	if _, err = client.TestTodo.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("test_todo: failed to feed test todo: %v", err)
	}
}

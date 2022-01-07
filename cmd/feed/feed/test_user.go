package feed

import (
	"context"
	"log"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/testuserprofile"
)

var testUserFeed = struct {
	bob ent.CreateTestUserInput
	ken ent.CreateTestUserInput
	tom ent.CreateTestUserInput
}{
	bob: ent.CreateTestUserInput{
		Name: "Bob",
		Age:  25,
		Profile: testuserprofile.TestUserProfile{
			Address: "address",
			Phone:   "09000000000",
			Body: testuserprofile.TestUserProfileBody{
				Weight: 60,
				Height: 180,
				Comment: testuserprofile.TestUserProfileBodyComment{
					Type: "paragraph",
					Text: "test",
				},
			}},
	},
	ken: ent.CreateTestUserInput{
		Name: "Ken",
		Age:  30,
		Profile: testuserprofile.TestUserProfile{
			Address: "address",
			Phone:   "09000000000",
			Body: testuserprofile.TestUserProfileBody{
				Weight: 60,
				Height: 180,
				Comment: testuserprofile.TestUserProfileBodyComment{
					Type: "paragraph",
					Text: "test",
				},
			}},
	},
	tom: ent.CreateTestUserInput{
		Name: "Tom",
		Age:  20,
		Profile: testuserprofile.TestUserProfile{
			Address: "address",
			Phone:   "09000000000",
			Body: testuserprofile.TestUserProfileBody{
				Weight: 60,
				Height: 180,
				Comment: testuserprofile.TestUserProfileBodyComment{
					Type: "paragraph",
					Text: "test",
				},
			}},
	},
}

// TestUser generates test user data
func TestUser(ctx context.Context, client *ent.Client) {
	_, err := client.TestUser.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("test_user: failed to delete test user: %v", err)
	}

	ts := []ent.CreateTestUserInput{
		testUserFeed.bob,
		testUserFeed.ken,
		testUserFeed.tom,
	}
	bulk := make([]*ent.TestUserCreate, len(ts))
	for i, t := range ts {
		bulk[i] = client.TestUser.Create().SetInput(t)
	}
	if _, err = client.TestUser.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("test_user: failed to feed test user: %v", err)
	}
}

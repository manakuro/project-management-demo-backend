package mutation_test

import (
	"net/http"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/infrastructure/router"
	"project-management-demo-backend/testutil"
	"project-management-demo-backend/testutil/e2e"
	"testing"
)

func TestTestUser_CreateTestUser(t *testing.T) {
	expect, teardown := e2e.Setup(t, e2e.SetupOption{
		Teardown: func(t *testing.T, client *ent.Client) {
			testutil.DropTestUser(t, client)
		},
	})
	defer teardown()

	r := expect.POST(router.QueryPath).WithJSON(map[string]string{
		"query": `
			mutation {  
				createTestUser(input: {name: "Tom1", age: 20}) {   
					age    
					name
					id    
					createdAt    
					updatedAt  
			}
		}`,
	}).Expect()

	r.Status(http.StatusOK)
	testUser := e2e.GetData(r, "createTestUser")

	testUser.Value("age").Number().Equal(20)
	testUser.Value("name").String().Equal("Tom1")
}

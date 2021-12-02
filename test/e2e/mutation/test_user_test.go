package mutation_test

import (
	"net/http"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/pkg/infrastructure/router"
	"project-management-demo-backend/testutil"
	"testing"
)

func TestTestUserMutation(t *testing.T) {
	expect, teardown := testutil.SetupE2E(t, testutil.SetupE2EOption{
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
	obj := r.JSON().
		Object().
		Value("data").
		Object().
		Value("createTestUser").
		Object()

	obj.Value("age").Number().Equal(20)
	obj.Value("name").String().Equal("Tom1")
}

package resolver

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/schema/ulid"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/controller"
	"sync"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Channels of subscription
type Channels struct {
	TestUserUpdated map[string]struct {
		id ulid.ID
		ch chan *ent.TestUser
	}
}

// Resolver is a context struct
type Resolver struct {
	client     *ent.Client
	controller controller.Controller
	channels   Channels

	mutex sync.Mutex
}

// NewSchema creates NewExecutableSchema
func NewSchema(client *ent.Client, controller controller.Controller) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client:     client,
			controller: controller,
			channels: Channels{
				TestUserUpdated: map[string]struct {
					id ulid.ID
					ch chan *ent.TestUser
				}{},
			},
			mutex: sync.Mutex{},
		},
	})
}

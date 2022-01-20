package resolver

import (
	"project-management-demo-backend/ent"
	"project-management-demo-backend/graph/generated"
	"project-management-demo-backend/pkg/adapter/controller"
	"project-management-demo-backend/pkg/util/subscription"
	"sync"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is a context struct.
type Resolver struct {
	client        *ent.Client
	controller    controller.Controller
	subscriptions *subscription.Subscriptions

	mutex sync.Mutex
}

// NewSchema creates NewExecutableSchema.
func NewSchema(client *ent.Client, controller controller.Controller) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client:     client,
			controller: controller,
			subscriptions: &subscription.Subscriptions{
				ColorUpdated:              map[string]subscription.ColorUpdated{},
				FavoriteProjectCreated:    map[string]subscription.FavoriteProjectCreated{},
				FavoriteProjectIDsUpdated: map[string]subscription.FavoriteProjectIDsUpdated{},
				IconUpdated:               map[string]subscription.IconUpdated{},
				MeUpdated:                 map[string]subscription.MeUpdated{},
				ProjectBaseColorUpdated:   map[string]subscription.ProjectBaseColorUpdated{},
				ProjectIconUpdated:        map[string]subscription.ProjectIconUpdated{},
				ProjectLightColorUpdated:  map[string]subscription.ProjectLightColorUpdated{},
				ProjectTeammateUpdated:    map[string]subscription.ProjectTeammateUpdated{},
				ProjectUpdated:            map[string]subscription.ProjectUpdated{},
				TeammateUpdated:           map[string]subscription.TeammateUpdated{},
				TestUserUpdated:           map[string]subscription.TestUserUpdated{},
				WorkspaceTeammateUpdated:  map[string]subscription.WorkspaceTeammateUpdated{},
				WorkspaceUpdated:          map[string]subscription.WorkspaceUpdated{},
			},
			mutex: sync.Mutex{},
		},
	})
}

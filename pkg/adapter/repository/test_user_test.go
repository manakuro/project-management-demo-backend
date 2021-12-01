package repository_test

import (
	"context"
	"project-management-demo-backend/config"
	"project-management-demo-backend/ent"
	"project-management-demo-backend/ent/enttest"
	"project-management-demo-backend/pkg/adapter/repository"
	"project-management-demo-backend/pkg/entity/model"
	"project-management-demo-backend/pkg/infrastructure/datastore"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"entgo.io/ent/dialect"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (client *ent.Client, teardown func()) {
	config.ReadConfig()

	d := datastore.New()
	c := enttest.Open(t, dialect.MySQL, d)

	return c, func() {
		ctx := context.Background()
		_, err := client.TestUser.Delete().Exec(ctx)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		_ = c.Close()
	}
}

type args struct {
	ctx context.Context
}

func Test_testUserRepository_List(t *testing.T) {
	t.Helper()

	client, teardown := setup(t)
	defer teardown()

	repo := repository.NewTestUserRepository(client)

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(ctx context.Context, t *testing.T) (us []*model.TestUser, err error)
		assert  func(t *testing.T, us []*model.TestUser, err error)
		args    struct {
			ctx context.Context
		}
		teardown func()
	}{
		{
			name: "It should get user's list",
			arrange: func(t *testing.T) {
				ctx := context.Background()
				_, err := client.TestUser.Delete().Exec(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}

				users := []struct {
					name string
					age  int
				}{
					{
						name: "test",
						age:  10,
					},
					{
						name: "test2",
						age:  11,
					},
					{
						name: "test3",
						age:  12,
					},
				}
				bulk := make([]*ent.TestUserCreate, len(users))
				for i, u := range users {
					bulk[i] = client.TestUser.Create().SetName(u.name).SetAge(u.age)
				}

				_, err = client.TestUser.
					CreateBulk(bulk...).
					Save(ctx)
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
			},
			act: func(ctx context.Context, t *testing.T) (us []*model.TestUser, err error) {
				return repo.List(ctx)
			},
			assert: func(t *testing.T, got []*model.TestUser, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 3, len(got))
			},
			args: args{
				ctx: context.Background(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange(t)
			got, err := tt.act(tt.args.ctx, t)
			tt.assert(t, got, err)
		})
	}
}

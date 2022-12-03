package postgres_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/oklog/ulid/v2"
	"github.com/pressly/goose/v3"
	"github.com/smartystreets/goconvey/convey"

	"chatty/chatty/app/config"
	"chatty/chatty/entity"
	repo "chatty/chatty/repository/postgres"
	"chatty/pkg/docker_test"
	"chatty/pkg/repository/postgres"
	_ "chatty/pkg/repository/postgres/migrations"
)

func TestPgUserRepo_AddUser(t *testing.T) {
	docker_test.InitEnv()
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		t.Error(err)
	}

	pgResource := docker_test.NewPostgres()
	defer func() {
		if err := pgResource.Pool.RemoveContainerByName(pgResource.Resource.Container.Name); err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(time.Second * 3)

	pgDb, err := postgres.Connection(ctx, cfg.Pg)
	if err != nil {
		t.Error(err)
	}

	RunPgMigrations(cfg.Pg)

	tests := []struct {
		name        string
		wantErr     bool
		expectEqual bool
		user        entity.User
	}{
		{
			name:        "Adding a user to the users table should pass without errors",
			wantErr:     false,
			expectEqual: true,
			user: entity.User{
				ID: entity.UserID(ulid.Make()),
				Creds: entity.UserCreds{
					Login:    "user1",
					Password: "user1Password",
				},
				Contacts: entity.UserContacts{
					Email:       "user1@testmail.com",
					PhoneNumber: "+71111111111",
				},
			},
		},
		{
			name:        "Adding a duplicated user should cause an error",
			wantErr:     true,
			expectEqual: false,
			user: entity.User{
				ID: entity.UserID(ulid.Make()),
				Creds: entity.UserCreds{
					Login:    "user1",
					Password: "user1Password",
				},
				Contacts: entity.UserContacts{
					Email:       "user1@testmail.com",
					PhoneNumber: "+71111111111",
				},
			},
		},
	}

	repo := repo.NewPgChattyRepo(pgDb)

	for _, tt := range tests {
		convey.Convey(tt.name, t, func() {
			err := repo.AddUser(ctx, tt.user)

			convey.So(err != nil, convey.ShouldEqual, tt.wantErr)

			if tt.expectEqual {
				user := entity.User{}
				row := pgDb.QueryRow(ctx, `select id, login, password, email, phone_number from public.user
							where login = $1`, tt.user.Creds.Login)
				err := row.Scan(&user.ID, &user.Creds.Login, &user.Creds.Password,
					&user.Contacts.Email, &user.Contacts.PhoneNumber)

				convey.So(err, convey.ShouldBeNil)
				convey.ShouldResemble(user, tt.user)
			}
		})
	}
}

func RunPgMigrations(cfgPg config.PG) error {
	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s port=%d sslmode=%s",
		cfgPg.User, cfgPg.Password, cfgPg.Host, cfgPg.DbName, cfgPg.Port, "disable")

	mdb, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = goose.Up(mdb, "/")
	if err != nil {
		return err
	}

	return nil
}

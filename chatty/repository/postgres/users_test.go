package postgres_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
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

	err = RunPgMigrations(cfg.Pg, int64(1))
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name         string
		wantErr      bool
		expectEqual  bool
		userCreds    entity.UserCreds
		userContacts entity.UserContacts
	}{
		{
			name:        "Adding a userCreds to the users table should pass without errors",
			wantErr:     false,
			expectEqual: true,
			userCreds: entity.UserCreds{
				Login:    "user1",
				Password: "user1Password",
			},
			userContacts: entity.UserContacts{
				Email:       "user1@testmail.com",
				PhoneNumber: "+71111111111",
			},
		},
		{
			name:        "Adding a duplicated userCreds should throw an error",
			wantErr:     true,
			expectEqual: false,
			userCreds: entity.UserCreds{
				Login:    "user1",
				Password: "user1Password",
			},
			userContacts: entity.UserContacts{
				Email:       "user1@testmail.com",
				PhoneNumber: "+71111111111",
			},
		},
	}

	repository := repo.NewPgChattyRepo(pgDb)

	for _, tt := range tests {
		convey.Convey(tt.name, t, func() {

			user := entity.NewUser(tt.userCreds, tt.userContacts)

			err := repository.AddUser(ctx, *user)

			convey.So(err != nil, convey.ShouldEqual, tt.wantErr)

			if tt.expectEqual {
				user := entity.User{}

				row := pgDb.QueryRow(ctx, `select id, login, password, email, phone_number from public.user
							where login = $1`, tt.userCreds.Login)
				err := row.Scan(&user.ID, &user.Creds.Login, &user.Creds.Password,
					&user.Contacts.Email, &user.Contacts.PhoneNumber)

				convey.So(err, convey.ShouldBeNil)
				convey.ShouldResemble(user, tt.userCreds)
			}
		})
	}
}

func TestPgUserRepo_GetUserByLogin(t *testing.T) {
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

	err = RunPgMigrations(cfg.Pg, int64(2))
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name        string
		wantErr     bool
		expectEqual bool
		userLogin   entity.UserLogin
	}{
		{
			name:        "Getting an existing userCreds from the users table should pass without errors",
			wantErr:     false,
			expectEqual: true,
			userLogin:   "user1",
		},
		{
			name:        "Getting a non-existent userCreds should throw an error",
			wantErr:     true,
			expectEqual: false,
			userLogin:   "userNotFoundLogin",
		},
	}

	repository := repo.NewPgChattyRepo(pgDb)

	for _, tt := range tests {
		convey.Convey(tt.name, t, func() {
			user, err := repository.GetUserByLogin(ctx, tt.userLogin)

			convey.So(err != nil, convey.ShouldEqual, tt.wantErr)

			if tt.expectEqual {
				convey.ShouldEqual(user.Creds.Login, tt.userLogin)
			}
		})
	}
}

func RunPgMigrations(cfgPg config.PG, v int64) error {
	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s port=%d sslmode=%s",
		cfgPg.User, cfgPg.Password, cfgPg.Host, cfgPg.DbName, cfgPg.Port, "disable")

	mdb, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = goose.UpTo(mdb, "/", v)
	if err != nil {
		return err
	}

	return nil
}

package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

const (
	AddUsersSqlQry = `
			insert into  public.user values
			('01GV0N3S3Y5JCV4EJSVD907AN2', 'user1', 'user1Password', 'user1@testmail.com', '+71111111111', now()),
			('01GV0N3S3Y5JCV4EJSVD907AN3', 'user2', 'user2Password', 'user2@testmail.com', '+71111111112', now()),
			('01GV0N3S3Y5JCV4EJSVD907AN4', 'user3', 'user3Password', 'user3@testmail.com', '+71111111113', now());
			`
	DeleteUsersSqlQry = `delete from public.user;`
)

func init() {
	goose.AddMigration(Up02, Down02)
}

func Up02(tx *sql.Tx) error {
	_, err := tx.Exec(AddUsersSqlQry)
	if err != nil {
		return err
	}

	return nil
}

func Down02(tx *sql.Tx) error {
	_, err := tx.Exec(DeleteUsersSqlQry)
	if err != nil {
		return err
	}

	return nil
}

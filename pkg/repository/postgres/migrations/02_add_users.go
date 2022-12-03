package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

const (
	AddUsersSqlQry = `
			insert into  public.user values
			('40e6215d-b5c6-4896-987c-f30f3678f608', 'user1', 'user1Password', 'user1@testmail.com', '+71111111111', now()),
			('40e6215d-b5c6-4896-987c-f30f3678f609', 'user2', 'user2Password', 'user2@testmail.com', '+71111111112', now()),
			('40e6215d-b5c6-4896-987c-f30f3678f610', 'user3', 'user3Password', 'user3@testmail.com', '+71111111113', now());
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

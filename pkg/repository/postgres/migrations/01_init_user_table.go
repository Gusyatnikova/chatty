package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

const (
	InitUserTableSqlQry = `create table if not exists public.user
				(
    				id           character(26)               primary key,
    				login        varchar                     not null unique,
    				password     varchar                     not null,
    				email        varchar                     not null unique,
    				phone_number varchar                     not null unique,
    				created      timestamp without time zone not null default now()
				);`
	DropUserTableSqlQry = `drop table public.user;`
)

func init() {
	goose.AddMigration(Up01, Down01)
}

func Up01(tx *sql.Tx) error {
	_, err := tx.Exec(InitUserTableSqlQry)
	if err != nil {
		return err
	}

	return nil
}

func Down01(tx *sql.Tx) error {
	_, err := tx.Exec(DropUserTableSqlQry)
	if err != nil {
		return err
	}

	return nil
}

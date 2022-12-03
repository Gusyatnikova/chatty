package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

const (
	InitUserTableSqlQry = `create table if not exists public.user
				(
					id           uuid                        primary key,
					login        varchar                     not null unique,
					password     varchar                     not null,
					email        varchar                     not null unique,
					phone_number varchar                     not null unique,
					created      timestamp without time zone not null default now()
				);`
	DropUserTableSqlQry = `drop table public.user;`
)

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sql.Tx) error {
	_, err := tx.Exec(InitUserTableSqlQry)
	if err != nil {
		return err
	}

	return nil
}

func Down00001(tx *sql.Tx) error {
	_, err := tx.Exec(DropUserTableSqlQry)
	if err != nil {
		return err
	}

	return nil
}

package postgres

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"chatty/chatty/app/config"
)

func Connection(pgConf config.PG) *pgxpool.Pool {
	//todo: use sprintf
	conf, err := pgxpool.ParseConfig("user=" + pgConf.User +
		" " + "password=" + pgConf.Password +
		" " + "host=" + pgConf.Host +
		" " + "dbname=" + pgConf.DbName +
		" " + "port=" + strconv.Itoa(pgConf.Port) +
		" " + "pool_max_conns=" + strconv.Itoa(pgConf.PoolMax))
	if err != nil {
		//todo: never use fatal, return error
		//add Conneection() (pool, error)
		log.Fatal(err)
	}

	//move tihs envs to config
	conf.HealthCheckPeriod = time.Second * 30
	conf.MaxConnIdleTime = time.Second
	conf.MinConns = 2

	pgConnect, err := pgxpool.ConnectConfig(context.Background(), conf)
	if err != nil {
		//
		log.Fatal(err)
	}

	return pgConnect
}

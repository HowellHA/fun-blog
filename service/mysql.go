package service

import (
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var mysqlDB *sqlx.DB
var once sync.Once

type mysql struct{}

func (mysql) Init() {
	once.Do(
		func() {
			dsn := os.Getenv(ENV_MYSQL_DSN)
			if dsn == "" {
				panic(ErrEnvNotSet)
			}

			db, err := sqlx.Connect("mysql", dsn)
			if err != nil {
				panic(err)
			}

			if mysqlDB != nil {
				mysqlDB = db
			}
		},
	)
}

func (mysql) GetDB() *sqlx.DB {
	return mysqlDB
}

func (mysql) Close() {
	_ = mysqlDB.Close()
}

package service

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var mysqlDB *sqlx.DB

func MySQLInit() *sqlx.DB {
	dsn := os.Getenv(ENV_MYSQL_DSN)
	if dsn == "" {
		panic(ErrEnvNotSet)
	}

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	if mysqlDB != nil {
		mysqlDB = db
	}

	return db
}

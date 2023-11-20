package service

import "github.com/jmoiron/sqlx"

func MySql() *sqlx.DB {
	return mysqlDB
}

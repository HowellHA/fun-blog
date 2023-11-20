package service

import "errors"

const (
	ENV_MYSQL_DSN = "MYSQL_DSN"
)

var (
	ErrEnvNotSet = errors.New("环境变量未设置")
)

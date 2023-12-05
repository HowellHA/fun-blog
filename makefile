dev: export MYSQL_DSN=fun:123456@tcp(127.0.0.1:3306)/test?timeout=30s&parseTime=true
dev:
	air -c .air.toml

lint:
	CGO_ENABLED=0 golangci-lint run --timeout=5m
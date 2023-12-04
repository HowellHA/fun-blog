dev:
	air -c .air.toml

lint:
	CGO_ENABLED=0 golangci-lint run --timeout=5m
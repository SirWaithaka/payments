-include .env
export

test:
	LOG_MODE=SILENT go test ./...

test.verbose:
	go test ./... -v

test.cover:
	go test ./... -v -coverprofile=coverage.out

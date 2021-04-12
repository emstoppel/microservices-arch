all: dependencies format imports mocking linter test
dependencies:
	@echo "Syncing dependencies with go mod tidy"
	@go mod tidy
format:
	@echo "Formatting Go code recursively"
	@go fmt ./...
imports:
	@echo "Executing goimports recursively"
	@goimports -w $(find . -type f -name '*.go') cmd/
mocking:
	@echo "generating mock files recursively"
	@go generate ./...
linter:
	@echo "Linting everything out"
	@golangci-lint run ./...
test:
	@echo "Running tests"
	@go test ./... -covermode=atomic -coverpkg=./... -count=1 -race
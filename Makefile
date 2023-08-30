
build-ast:
	@go run ./cmd/astgen/astgen.go
	@gofmt -w ./ast

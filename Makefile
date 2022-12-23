
.PHONY: run
run:
	go run cmd/server/main.go
.PHONY: test
test:
	go test -race -v ./...

.PHONY: generate
generate:
	go generate ./...
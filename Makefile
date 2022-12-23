
.PHONY: test
test:
	go test -race -v ./...

.PHONY: generate
generate:
	go generate ./...
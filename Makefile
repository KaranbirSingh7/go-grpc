APPLICATION_NAME=go-grpc
DOCKER_USERNAME=karanbirsingh
_BUILD_ARGS_TAG ?= ${GIT_HASH}
_BUILD_ARGS_RELEASE_TAG ?= latest
_BUILD_ARGS_DOCKERFILE ?= Dockerfile
GIT_HASH ?= $(shell git log --format="%h" -n 1)


.PHONY: fmt
fmt:
	@gofmt -l -w -s .

.PHONY: build
build:
	rm -rf bin/ && go build -o bin/app cmd/server/main.go
	echo "binary -> bin/app"

.PHONY: run
run:
	go run cmd/server/main.go
.PHONY: test
test:
	go test -race -v ./...

.PHONY: generate
generate:
	go generate ./...

.PHONY: docker-build
docker-build:
	docker build --tag ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG} -f ${_BUILD_ARGS_DOCKERFILE} .


.PHONY: docker-run
docker-run: docker-build
	docker run ${DOCKER_USERNAME}/${APPLICATION_NAME}:${_BUILD_ARGS_TAG}

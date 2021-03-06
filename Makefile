.PHONY: all clean echo test fmt install bench run bootstrap build docker-build

GDFLAGS ?= $(GDFLAGS:)
ARGS ?= $(ARGS:)
TAG ?= latest

EXTERNAL_TOOLS=\
	github.com/tools/godep \
	golang.org/x/tools/cmd/vet

all: test

bootstrap:
	@for tool in  $(EXTERNAL_TOOLS) ; do \
		echo "===> Installing $$tool" ; \
    go get $$tool; \
	done

clean:
	@echo "===> Cleaning"
	@godep go clean $(GDFLAGS) -i ./...

build: fmt lint
	@echo "===> Building"
	@godep go build $(GDFLAGS)  ./...

fmt:
	@echo "===> Formatting"
	@godep go fmt $(GDFLAGS) ./...

lint:
	@echo "===> Linting with vet"
	@godep go vet $(GDFLAGS) ./...

install: build
	@echo "===> Installing"
	@godep go install $(GDFLAGS)

test:
	@echo "===> Testing"
	@godep go test $(GDFLAGS) ./...

bench:
	@echo "===> Benchmarking"
	@godep go test -run=NONE -bench=. $(GDFLAGS) ./...

start: build
	@echo "===> Starting Server"
	@./$(EXECUTABLE) $(ARGS)

run:
	@echo "===> Running Server"
	@godep go run *.go

docker-build:
	@echo "===> Building Dockerfile"
	@eval "$(docker-machine env)"
	@docker build -t jalateras/version .

docker-push:
	@echo "===> Pushing Image"
	@eval "$(docker-machine env)"
	@docker push jalateras/version:$(TAG)

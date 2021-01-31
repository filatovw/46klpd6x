APP := api

DC = docker-compose

REVISION ?= $(shell git rev-parse --short=8 HEAD)
VERSION ?= $(shell cat ./VERSION)
BUILDTIME = $(shell date -u +"%Y%m%d%H%M%S")
VERSION_REVISION ?= ${VERSION}-${REVISION}
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)

DOCKERHUB :=
DOCKER_DIR := docker/
GO_IMAGE := golang:1.15.7

BIN_DIR ?= bin/
STORAGE_DIR ?= storage/


##    clean - delete everything
.PHONY: clean
clean:
	@rm -f BIN_DIR/*
	$(DC) down


##    api/build - build api service
.PHONY: api/build
api/build:
	go build -o $(BIN_DIR)$(APP) ./cmd/$(APP)

##    api/test - test api service
api/test:
	go test -v -race ./...

##    api/lint - run linter
api/lint:
	go vet ./...

##    api/up - up service
.PHONY: api/up
api/up: api/build
	$(BIN_DIR)$(APP)

##    docker/api/build - build API service container
.PHONY: docker/api/build
docker/api/build:
	docker build \
		--no-cache \
		--build-arg REVISION=$(REVISION) \
		--build-arg BUILDTIME=$(BUILDTIME) \
		--build-arg VERSION=$(VERSION) \
		--build-arg APP=$(APP) \
		--rm -f $(DOCKER_DIR)$(APP)/Dockerfile \
		-t $(DOCKERHUB)$(APP):latest \
		-t $(DOCKERHUB)$(APP):$(VERSION) \
		-t $(DOCKERHUB)$(APP):$(BRANCH) \
		-t $(DOCKERHUB)$(APP):$(VERSION_REVISION) .

##    docker/api/up - run API service in docker
docker/api/up:
	$(DC) up -d $(APP)

##    docker/api/down - stop running API container
docker/api/down:
	$(DC) stop $(APP)

##    docker/infra/up - run API service in docker
docker/infra/up:
	$(DC) up -d db db-migrate

##    docker/infra/down - stop running API container
docker/infra/down:
	$(DC) stop db

##    docker/logs - watch logs
docker/logs:
	$(DC) logs -f

##    docker/<local target> - universal docker wrapper for local targets
docker/%:
	docker run --rm \
		-v $(PWD)/:/usr/app/ \
		-w /usr/app/ \
		-e APP=$(APP) \
		${GO_IMAGE} make $*

##    migration/new/<migration name> - create new migration file
migration/new/%:
	docker run --rm -it \
		-v $(PWD)/migrations:/db/migrations --network=host \
		amacneil/dbmate new $*

##    migration/up - apply db migration
migration/up:
	$(DC) up db-migrate

##    migration/down - rollback db migration
migration/down:
	$(DC) up db-rollback

##    migration/drop - drop database
migration/drop:
	$(DC) up db-drop



## -
## help - this message
.PHONY: help
help: Makefile
	@echo "Application: ${APP}\n"
	@echo "Run command:\n  make <target>\n"
	@grep -E -h '^## .*' $(MAKEFILE_LIST) | sed -n 's/^##//p'  | column -t -s '-' |  sed -e 's/^/ /'

APP := api

REVISION ?= $(shell git rev-parse --short=8 HEAD)
VERSION ?= $(shell cat ./VERSION)
BUILDTIME = $(shell date -u +"%Y%m%d%H%M%S")
VERSION_REVISION ?= ${VERSION}-${REVISION}
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
IMAGE := $(REPO):$(VERSION)

DOCKERHUB :=
DEPLOYMENT_DIR := deployment/

BIN_DIR ?= bin/

##    api/build - build api service
.PHONY: api/build
api/build:
	go build -o $(BIN_DIR)$(APP) ./cmd/$(APP)

##    api/up - up service
.PHONY: api/up
api/up: api/build
	$(BIN_DIR)$(APP)

##    docker/api/build - build api service in docker
.PHONY: docker/api/build
docker/api/build:
	docker build \
		--no-cache \
		--build-arg REVISION=$(REVISION) \
		--build-arg BUILDTIME=$(BUILDTIME) \
		--build-arg VERSION=$(VERSION) \
		--build-arg APP=$(APP) \
		--rm -f $(DEPLOYMENT_DIR)$(APP)/Dockerfile \
		-t $(DOCKERHUB)$(APP):latest \
		-t $(DOCKERHUB)$(APP):$(VERSION) \
		-t $(DOCKERHUB)$(APP):$(BRANCH) \
		-t $(DOCKERHUB)$(APP):$(VERSION_REVISION) .

##    docker/api/up - up service in docker
docker/api/up:
	docker-compose up -d $(APP)

##    docker/api/down - stop running API container
docker/api/down:
	docker-compose stop $(APP)

##    docker/logs - watch logs
docker/logs:
	docker-compose logs -f


## -
## help - this message
.PHONY: help
help: Makefile
	@echo "Application: ${APP}\n"
	@echo "Run command:\n  make <target>\n"
	@grep -E -h '^## .*' $(MAKEFILE_LIST) | sed -n 's/^##//p'  | column -t -s '-' |  sed -e 's/^/ /'

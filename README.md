# Project: 46klpd6x

# Code challenge

Read about the details [here](docs/CODE_CHALLENGE.md)

# Architecture

HTTP API service
    - Authentication is done using JWT token
    - Postgresql is used as a persistent storage
    - Redis is a temporary storage to keep cache

# Execution

## pre-requisites:

- docker
- docker-compose
- maketool

Start service:

    make api/up

Stop running service

    make api/down

Start infrastructure:

    make infra/up

Stop running service and shutdown the infrastructure:

    make infra/down

Apply migrations

    make db/migrate

Shutdown running services, cleanup the environment

    make all/clean

# Development

## pre-requisites:

- go >=v1.6
- maketool
- pre-commit


Build service locally

    make api/build

Run tests locally

    make api/test

Build from scratch and run

    make all

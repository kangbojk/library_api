.PHONY: all

build-docker: dependencies build-api-docker
build: dependencies build-api

all:
	build-docker

dependencies:
	go mod download

build-api-docker:
	go build -tags docker -o bin/main cmd/main.go

build-api:
	go build -tags local -o bin/main cmd/main.go
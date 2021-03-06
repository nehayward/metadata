SHELL:=/bin/bash
VERSION?=$(shell git describe --tags --always)
CURRENT_DOCKER_IMAGE=nehayward/metadata:$(VERSION)
LATEST_DOCKER_IMAGE=nehayward/metadata:latest

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --ldflags "-X main.Version=$(VERSION)" -o metadata .
	docker build -t $(CURRENT_DOCKER_IMAGE) .

test:
	go test ./... -v

run:
	go run main.go

release: build
	docker push $(CURRENT_DOCKER_IMAGE)
	docker tag  $(CURRENT_DOCKER_IMAGE) $(LATEST_DOCKER_IMAGE)
	docker push $(LATEST_DOCKER_IMAGE)

.PHONY: build test run release
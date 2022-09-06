#!/bin/bash
SHELL := /bin/bash

run:
	@go build && ./app

run-docker:
	@go build && ./app -env=docker

docker-up:
	@docker-compose --env-file=./.Docker/.docker.env up -d

docker-image:
	@docker build -t goapp .

proto-gen:
	@protoc module/articles/handler/grpc/*.proto --proto_path=. --go_out=. --go-grpc_out=.
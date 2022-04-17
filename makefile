#!/bin/bash
SHELL := /bin/bash

run:
	@go build && ./app

run-docker:
	@go build && ./app -env=docker

docker-up:
	@docker-compose --env-file=./env/.docker.env up -d

docker-down:
	@docker-compose down
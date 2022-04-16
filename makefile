#!/bin/bash
SHELL := /bin/bash

run:
	@go build && ./app

run-docker:
	@go build && ./app -env=docker
# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
THIS_FILE := $(lastword $(MAKEFILE_LIST))

APP = yacm

.PHONY: run-all-tests

build:
	go build -o $(APP) .

run:
	go run .

test:
	go test -cover ../../...

# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
THIS_FILE := $(lastword $(MAKEFILE_LIST))

.PHONY: test

test:
	@$(MAKE) -f $(THIS_FILE) \
		test-runit && \
	$(MAKE) -f $(THIS_FILE) \
		test-ubuntu;

test-build-base:
	@docker build \
		-t yacm-test-base .;

test-runit:
	@$(MAKE) -f $(THIS_FILE) \
		test-build-base && \
	docker build  \
		-t yacm-artix-runit \
		test/artix-runit/ && \
	docker run \
		--rm \
		--name yacm-artix-runit-test \
		yacm-artix-runit;

test-ubuntu:
	@$(MAKE) -f $(THIS_FILE) \
		test-build-base && \
	docker build  \
		-t yacm-ubuntu \
		test/ubuntu/ && \
	docker run \
		--rm \
		--name yacm-ubuntu-test \
		yacm-ubuntu;

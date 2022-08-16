# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
THIS_FILE := $(lastword $(MAKEFILE_LIST))

.PHONY: test

test:
	@$(MAKE) -f $(THIS_FILE) \
		test-runit;

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
		--name yacm-runit-test \
		yacm-runit;

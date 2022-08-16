# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
THIS_FILE := $(lastword $(MAKEFILE_LIST))

.PHONY: test

test:
	@$(MAKE) -f $(THIS_FILE) \
		test-runit && \
	@$(MAKE) -f $(THIS_FILE) \
		test-pacman && \
	@$(MAKE) -f $(THIS_FILE) \
		test-paru && \
	$(MAKE) -f $(THIS_FILE) \
		test-apt;

test-build-base:
	@docker build \
		-t yacm-base-test .;

test-runit:
	@$(MAKE) -f $(THIS_FILE) \
		test-build-base && \
	docker build  \
		-t yacm-runit \
		test/runit/ && \
	docker run \
		--rm \
		--name yacm-runit-test \
		yacm-runit;

test-apt:
	@$(MAKE) -f $(THIS_FILE) \
		test-build-base && \
	docker build  \
		-t yacm-apt \
		test/apt/ && \
	docker run \
		--rm \
		--name yacm-apt-test \
		yacm-apt;

test-pacman:
	@$(MAKE) -f $(THIS_FILE) \
		test-build-base && \
	docker build  \
		-t yacm-pacman \
		test/pacman/ && \
	docker run \
		--rm \
		--name yacm-pacman-test \
		yacm-pacman;

test-paru:
	@$(MAKE) -f $(THIS_FILE) \
		test-build-base && \
	docker build  \
		-t yacm-paru \
		test/paru/ && \
	docker run \
		--rm \
		--name yacm-paru-test \
		yacm-paru;

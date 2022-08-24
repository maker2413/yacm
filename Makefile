# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
THIS_FILE := $(lastword $(MAKEFILE_LIST))

.PHONY: run-all-tests

run-all-tests:
	@$(MAKE) -f $(THIS_FILE) \
		TEST=runit run-tests && \
	$(MAKE) -f $(THIS_FILE) \
		TEST=apt run-tests && \
	$(MAKE) -f $(THIS_FILE) \
		TEST=dnf run-tests && \
	$(MAKE) -f $(THIS_FILE) \
		TEST=pacman run-tests && \
	$(MAKE) -f $(THIS_FILE) \
		TEST=paru run-tests && \
	$(MAKE) -f $(THIS_FILE) \
		TEST=yay run-tests && \
	$(MAKE) -f $(THIS_FILE) \
		TEST=yum run-tests;

build-test-base:
	@docker build \
		-t yacm-base .;

run-tests:
	@$(MAKE) -f $(THIS_FILE) \
		build-test-base && \
	docker build \
		-t yacm-$(TEST) \
		test/$(TEST)/ && \
	docker run \
		--rm \
		--name yacm-$(TEST)-test \
		yacm-$(TEST);

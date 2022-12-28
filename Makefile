# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
THIS_FILE := $(lastword $(MAKEFILE_LIST))

.PHONY: run-all-tests

run-all-tests:
	@$(MAKE) -f $(THIS_FILE) \
		run-all-system-tests;

run-all-prompt-tests:
	@$(MAKE) -f $(THIS_FILE) \
		TEST=list run-prompt-tests;

run-prompt-tests:
	@$(MAKE) -f $(THIS_FILE) \
		build-test-base && \
	./test/run-prompt-tests.sh;

run-all-system-tests:
	@$(MAKE) -f $(THIS_FILE) \
		TEST=systemd run-system-tests && \
	$(MAKE) -f $(THIS_FILE) \
		TEST=runit run-system-tests && \
	$(MAKE) -f $(THIS_FILE) \
		TEST=apt run-system-tests && \
	$(MAKE) -f $(THIS_FILE) \
		TEST=dnf run-system-tests && \
	$(MAKE) -f $(THIS_FILE) \
		TEST=pacman run-system-tests && \
	$(MAKE) -f $(THIS_FILE) \
		TEST=paru run-system-tests && \
	$(MAKE) -f $(THIS_FILE) \
		TEST=yay run-system-tests && \
	$(MAKE) -f $(THIS_FILE) \
		TEST=yum run-system-tests;

run-system-tests:
	@$(MAKE) -f $(THIS_FILE) \
		build-test-base && \
	./test/run-system-tests.sh;

build-test-base:
	@podman build \
		-t yacm-base .;

#!/bin/bash

set -e;

cd "${0%/*}";

podman build \
       -t yacm-$TEST-prompt ./prompt-tests/$TEST;

#!/bin/bash

set -e;

cd "${0%/*}";

docker build \
       -t yacm-$TEST ./$TEST;

docker run -d \
       --tmpfs /tmp \
       --tmpfs /run \
       -v /sys/fs/cgroup:/sys/fs/cgroup:ro \
       -p 9090:80 \
       --name yacm-$TEST-test \
       --rm \
       yacm-$TEST;

docker exec -it \
       yacm-$TEST-test \
       /root/run-yacm.sh;

curl localhost:9090;

docker stop yacm-$TEST-test;

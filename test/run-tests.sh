#!/bin/bash

set -e;

cd "${0%/*}";

: "${DOCKER_PORT:=80}"

docker build \
       -t yacm-$TEST ./$TEST;

docker run -d \
       --tmpfs /tmp \
       --tmpfs /run \
       -v /sys/fs/cgroup:/sys/fs/cgroup:ro \
       -p 9090:$DOCKER_PORT \
       --name yacm-$TEST-test \
       --rm \
       yacm-$TEST;

docker exec -it \
       yacm-$TEST-test \
       /root/run-yacm.sh;

echo "Waiting for 2 seconds to give httpd time to boot";
sleep 2;

curl localhost:9090 || docker stop yacm-$TEST-test;

docker stop yacm-$TEST-test;

#!/bin/bash

set -e;

cd "${0%/*}";

: "${CONTAINER_PORT:=80}"

podman build \
       -t yacm-$TEST ./$TEST;

podman run -d \
       -p 9090:$CONTAINER_PORT \
       --name yacm-$TEST-test \
       --rm \
       yacm-$TEST;

case "$TEST" in
  systemd|runit)
    podman exec -it \
           yacm-$TEST-test \
           /root/run-yacm.sh;

    echo "Waiting for 2 seconds to give httpd time to boot";
    sleep 2;

    curl localhost:9090 || podman stop yacm-$TEST-test;;
  apt)
    echo "Waiting for 20 seconds to give httpd time to boot";
    sleep 20;

    podman exec -it \
           yacm-$TEST-test \
           /bin/bash -c "which apache2" \
      || podman stop yacm-$TEST-test;;
esac

podman stop yacm-$TEST-test;

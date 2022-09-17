#!/bin/bash

set -e;

cd "${0%/*}";

: "${RUNIT_PORT:=8080}"
: "${SYSTEMD_PORT:=80}"
: "${TEST:=runit}"

function wait_for_bootstrap() {
  local __package=$1;
  local found=0;
  local timer=0;

  while [[ "$found" == 0 && "$timer" < 600 ]]; do
    echo "Waiting for bootstrap to complete...";

    sleep 5;

    if podman exec -it -e package=$__package yacm-$TEST-test /bin/bash -c 'which $package'; then
      found=1;
    fi

    timer+=5;
  done
}

podman build \
       -t yacm-$TEST ./$TEST;

case "$TEST" in
  systemd)
    podman run -d \
           -p 9090:$SYSTEMD_PORT \
           --name yacm-$TEST-test \
           --rm \
           yacm-$TEST;;
  runit)
    podman run -d \
           -p 9090:$RUNIT_PORT \
           --name yacm-$TEST-test \
           --rm \
           yacm-$TEST;;
  *)
    podman run -d \
            --name yacm-$TEST-test \
            --rm \
            yacm-$TEST;;
esac

case "$TEST" in
  systemd|runit)
    podman exec -it \
           yacm-$TEST-test \
           /root/run-yacm.sh;

    echo -e "\nWaiting for 10 seconds to give httpd time to boot";
    sleep 10;

    curl localhost:9090 || podman stop yacm-$TEST-test;;
  apt)
    wait_for_bootstrap apache2;

    podman exec -it \
           yacm-$TEST-test \
           /bin/bash -c "which apache2" \
      || podman stop yacm-$TEST-test;;
  dnf|yum)
    wait_for_bootstrap httpd;

    podman exec -it \
           yacm-$TEST-test \
           /bin/bash -c "which httpd" \
      || podman stop yacm-$TEST-test;;
  pacman|paru|yay)
    wait_for_bootstrap lighttpd;

    podman exec -it \
           yacm-$TEST-test \
           /bin/bash -c "which lighttpd" \
      || podman stop yacm-$TEST-test;;
esac

podman stop yacm-$TEST-test;

FROM yacm-base
FROM docker.io/fedora:latest

COPY run-yacm.sh /root/

COPY dnf-profile.yml /root/

COPY wait_for_tests.sh /root/

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

CMD /root/run-yacm.sh

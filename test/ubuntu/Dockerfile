FROM yacm-test-base
FROM ubuntu

RUN apt-get update && \
  apt-get install -y --no-install-recommends \
    systemd systemd-sysv dbus dbus-user-session

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

COPY run-tests.sh /root/

COPY ubuntu-profile.yml /root/

CMD ["/usr/sbin/initctl", "/root/run-tests.sh"]

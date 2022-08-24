FROM yacm-base
FROM fedora

RUN dnf install -y systemd procps less

COPY run-yacm.sh /root/

COPY systemd-profile.yml /root/

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

CMD [ "/sbin/init" ]

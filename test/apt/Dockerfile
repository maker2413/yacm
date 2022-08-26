FROM yacm-base
FROM ubuntu

RUN apt update && apt install -y systemd procps less

RUN apt install -y openssh-server

RUN systemctl enable ssh

COPY run-yacm.sh /root/

COPY apt-profile.yml /root/

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

CMD [ "/sbin/init" ]

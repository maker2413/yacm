FROM yacm-base
FROM docker.io/ubuntu:latest

RUN apt update && apt install -y systemd procps less

COPY run-yacm.sh /root/

COPY apt-profile.yml /root/

COPY wait_for_tests.sh /root/

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

CMD [ "/root/run-yacm.sh" ]

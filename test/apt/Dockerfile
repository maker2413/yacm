FROM yacm-base-test
FROM ubuntu

COPY run-tests.sh /root/

COPY apt-profile.yml /root/

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

CMD /root/run-tests.sh

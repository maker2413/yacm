FROM yacm-base-test
FROM ubuntu

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

COPY run-tests.sh /root/

COPY apt-profile.yml /root/

CMD /root/run-tests.sh

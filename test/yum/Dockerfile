FROM yacm-base-test
FROM fedora

COPY run-tests.sh /root/

COPY yum-profile.yml /root/

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

CMD /root/run-tests.sh

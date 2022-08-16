FROM yacm-base-test
FROM artixlinux/runit

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

COPY run-tests.sh /root/

COPY runit-profile.yml /root/

CMD /root/run-tests.sh

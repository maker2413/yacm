FROM yacm-base
FROM artixlinux/runit

COPY run-tests.sh /root/

COPY runit-profile.yml /root/

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

CMD /root/run-tests.sh

FROM yacm-base-test
FROM archlinux

COPY run-tests.sh /root/

COPY pacman-profile.yml /root/

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

CMD /root/run-tests.sh

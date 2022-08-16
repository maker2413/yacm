FROM yacm-base-test
FROM archlinux

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

COPY run-tests.sh /root/

COPY pacman-profile.yml /root/

CMD /root/run-tests.sh

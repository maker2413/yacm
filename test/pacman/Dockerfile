FROM yacm-base
FROM archlinux

COPY run-yacm.sh /root/

COPY pacman-profile.yml /root/

COPY wait_for_tests.sh /root/

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

CMD /root/run-yacm.sh

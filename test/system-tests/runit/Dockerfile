FROM yacm-base
FROM docker.io/artixlinux/artixlinux:base-runit

COPY run-yacm.sh /root/

COPY runit-profile.yml /root/

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

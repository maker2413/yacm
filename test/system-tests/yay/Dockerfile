FROM yacm-base
FROM docker.io/archlinux:latest

RUN pacman -Sy sudo --needed --noconfirm

RUN useradd -mG wheel yaytest

RUN sed -Ei 's/^#\ (%wheel.*NOPASSWD.*)/\1/' /etc/sudoers

COPY run-yacm.sh /home/yaytest/

COPY yay-profile.yml /home/yaytest/

COPY install-yay.sh /home/yaytest/

COPY wait_for_tests.sh /home/yaytest/

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

USER yaytest

WORKDIR /home/yaytest

CMD ./run-yacm.sh

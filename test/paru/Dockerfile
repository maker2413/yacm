FROM yacm-base
FROM docker.io/archlinux:latest

RUN pacman -Sy sudo --needed --noconfirm

RUN useradd -mG wheel parutest

RUN sed -Ei 's/^#\ (%wheel.*NOPASSWD.*)/\1/' /etc/sudoers

COPY run-yacm.sh /home/parutest/

COPY paru-profile.yml /home/parutest/

COPY install-paru.sh /home/parutest/

COPY wait_for_tests.sh /home/parutest/

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

USER parutest

WORKDIR /home/parutest

CMD ./run-yacm.sh

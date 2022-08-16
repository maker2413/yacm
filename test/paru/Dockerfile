FROM yacm-base-test
FROM archlinux

RUN pacman -Sy sudo --needed --noconfirm

RUN useradd -mG wheel parutest

RUN sed -Ei 's/^#\ (%wheel.*NOPASSWD.*)/\1/' /etc/sudoers

COPY run-tests.sh /home/parutest/

COPY paru-profile.yml /home/parutest/

COPY install-paru.sh /home/parutest/

COPY --from=0 /usr/bin/yacm /usr/bin/yacm

USER parutest

WORKDIR /home/parutest

CMD ./run-tests.sh

FROM yacm-base
FROM docker.io/archlinux:latest

RUN pacman -Sy sudo --needed --noconfirm

RUN useradd -mG wheel listtest

RUN sed -Ei 's/^#\ (%wheel.*NOPASSWD.*)/\1/' /etc/sudoers

USER listtest

WORKDIR /home/listtest

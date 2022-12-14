#!/bin/sh

which paru > /dev/null 2>&1;

if [ $? -ne 0 ];
then
  sudo pacman -Sy git base-devel --needed --noconfirm;
  git clone https://aur.archlinux.org/paru-bin.git;
  cd paru-bin || exit;
  makepkg -si --noconfirm;
  cd ..;
  rm -rf paru-bin;
fi

paru -S paru-bin --needed --noconfirm --skipreview;

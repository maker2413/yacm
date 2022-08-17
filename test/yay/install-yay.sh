#!/bin/sh

which yay > /dev/null 2>&1;

if [ $? -ne 0 ];
then
  sudo pacman -Sy git base-devel --needed --noconfirm;
  git clone https://aur.archlinux.org/yay-bin.git;
  cd yay-bin;
  makepkg -si --noconfirm;
  cd ..;
  rm -rf yay-bin;
fi

yay -S yay-bin --needed --noconfirm;

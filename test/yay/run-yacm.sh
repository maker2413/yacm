#!/bin/bash

yacm --yacm-base-dir ~/ \
     --yacm-dir ~/ \
     --yacm-profiles-dir ~/ \
     --yacm-scripts-dir ~/ \
     --skip-confirmation \
     bootstrap profile yay-profile;

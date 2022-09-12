#!/bin/bash

mkdir -p /run/runit/service

yacm --yacm-base-dir /root/ \
     --yacm-dir /root/ \
     --yacm-profiles-dir /root/ \
     --yacm-scripts-dir /root/ \
     --skip-confirmation \
     bootstrap profile runit-profile;

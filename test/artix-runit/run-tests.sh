#!/bin/bash

mkdir -p /run/runit/service

yacm --yacm-base-dir /root/ --yacm-dir /root/ --yacm-profiles-dir /root/ \
     bootstrap profile artix-runit-profile;

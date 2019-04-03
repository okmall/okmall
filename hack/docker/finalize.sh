#!/bin/sh
set -x
set -e

# Create user for OkMall
addgroup -S okmall
adduser -G okmall -H -D -g 'OkMall User' okmall -h /data/okmall -s /bin/bash && usermod -p '*' okmall && passwd -u okmall
echo "export OKMALL_CUSTOM=${OKMALL_CUSTOM}" >> /etc/profile

# Final cleaning
rm /app/okmall/docker/finalize.sh
rm /app/okmall/docker/nsswitch.conf

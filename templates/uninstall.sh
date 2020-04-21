#!/usr/bin/env bash

echo "Uninstalling virusstracker api"

systemctl stop virustracker.service
systemctl disable virustracker.service
systemctl daemon-reload

export VIRUSTRACKER_HOME=/opt/cs50vn/virustracker

rm -rf $VIRUSTRACKER_HOME
rm /etc/systemd/system/virustracker.service

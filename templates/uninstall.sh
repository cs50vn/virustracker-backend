#!/usr/bin/env bash

echo "Uninstalling virusstracker backend"

systemctl stop virustracker-backend.service
systemctl disable virustracker-backend.service
systemctl daemon-reload

export VIRUSTRACKER_HOME=/opt/cs50vn/virustracker-backend

rm -rf $VIRUSTRACKER_HOME
rm /etc/systemd/systemvirustracker-backend.service

#!/usr/bin/env bash

echo "Install virustracker backend service"

export VIRUSTRACKER_HOME=/opt/cs50vn/virustracker-backend

mkdir -p $VIRUSTRACKER_HOME

cp virustracker-backend config.json virustracker-backend.db $VIRUSTRACKER_HOME
cp virustracker-backend.service /etc/systemd/system

systemctl enable virustracker-backend.service
systemctl daemon-reload
systemctl start virustracker-backend.service

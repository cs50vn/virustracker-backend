#!/usr/bin/env bash


echo "Install virustracker service"

export VIRUSTRACKER_HOME=/opt/cs50vn/virustracker

mkdir  $VIRUSTRACKER_HOME

cp virustracker config.json virustracker.db $VIRUSTRACKER_HOME
cp virustracker.service /etc/systemd/system

#systemctl enable virustracker.service
#systemctl daemon-reload
#systemctl start virustracker.service

#!/usr/bin/env bash

echo "Starting before-remove.sh script"

if [ $(systemctl is-active pulsard.service) == active ]; then
  echo "Stopping pulsard"
  systemctl stop pulsard.service
fi

if [ $(systemctl is-active insolard.service) == active ]; then
  echo "Stopping insolard"
  systemctl stop insolard.service
fi

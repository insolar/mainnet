#!/usr/bin/env bash

echo "Starting before-install.sh script"

dirs=(
  /opt/insolar
  /opt/insolar/data
  /opt/insolar/config
  /opt/insolar/pulsar
  /opt/insolar/backup
  /opt/insolar/backup/tmp
  /opt/insolar/backup/target
  /opt/insolar/backup/received
  /opt/insolar/backup/merged
  /etc/insolar
)

useradd -m -s /bin/bash insolar && echo "User insolar created"

for dir in "${dirs[@]}"; do
  mkdir -p "$dir" && echo "Directory $dir created"
  chown insolar:insolar "$dir"
  chmod 0750 "$dir"
done

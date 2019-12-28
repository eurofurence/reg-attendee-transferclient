#! /bin/bash

set -o errexit

if [[ "$RUNTIME_USER" == "" ]]; then
  echo "RUNTIME_USER not set, bailing out. Please run setup.sh first."
  exit 1
fi

mkdir -p tmp
cp attendee-transfer-client tmp/
cp config.yaml tmp/
cp run-attendee-transfer.sh tmp/

chgrp $RUNTIME_USER tmp/*
chmod 640 tmp/config.yaml
chmod 750 tmp/attendee-transfer-client
chmod 750 tmp/run-attendee-transfer.sh
mv tmp/attendee-transfer-client /home/$RUNTIME_USER/work/attendee-transfer-client/
mv tmp/config.yaml /home/$RUNTIME_USER/work/attendee-transfer-client/
mv tmp/run-attendee-transfer.sh /home/$RUNTIME_USER/work/


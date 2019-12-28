#! /bin/bash

STARTTIME=$(date '+%Y-%m-%d_%H-%M-%S')

echo "Writing log to ~/work/logs/attendee-transfer-client.$STARTTIME.log"

cd ~/work/attendee-transfer-client

./attendee-transfer-client &> ~/work/logs/attendee-transfer-client.$STARTTIME.log


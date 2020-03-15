#!/bin/bash
while true
do
  echo "[loop.sh] Starting..."
  ./sungrabber
  if [ $? -eq 0 ]
  then
    echo "[loop.sh] Restarting."
  else
    echo "[loop.sh] Error exit code. Exiting."
    exit 1
  fi
done

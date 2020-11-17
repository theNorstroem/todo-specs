#!/usr/bin/env bash
# exit when any command fails
set -e

CLIENTPROJECTDIR="../todo-client"

if [ -z "$CLIENTPROJECTDIR" ]
then
      echo "$CLIENTPROJECTDIR is not set in your Env"
      exit 1
else
      echo "copy env to $CLIENTPROJECTDIR/src/configs/data_environment.js"
      cp ./dist/env.js "$CLIENTPROJECTDIR/src/configs/data_environment.js"
fi
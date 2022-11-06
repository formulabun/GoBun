#!/bin/bash

cd /usr/games/SRB2Kart || exit

DIRSTRUCTURE=$(find /addons -type d -regex "^.*/[0-9]+[^/]*$")
ADDONS=$(find /addons -type f,l | sort)

if [ -n "$DIRSTRUCTURE" ]; then
  echo "Using directory order as load order"
  /usr/bin/srb2kart -dedicated $* -file $ADDONS
else
  # Intentional word splitting
  echo "Using command line arguments as load order"
  /usr/bin/srb2kart -dedicated $*
fi

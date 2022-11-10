#!/bin/bash

cd /usr/games/SRB2Kart || exit

DIRSTRUCTURE=$(find /addons -type d -regex "^.*/[0-9]+[^/]*$")
ADDONS=$(find /addons -type f | sort)

if [ -n "$DIRSTRUCTURE" ]; then
  echo "srb2kart Docker | Using directory order as load order"
  /usr/bin/srb2kart -dedicated -file $ADDONS $*
else
  echo "srb2kart Docker | Using command line arguments as load order"
  /usr/bin/srb2kart -dedicated $*
fi

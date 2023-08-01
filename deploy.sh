#!/bin/bash

MY_DIR="$(dirname "$(readlink -e "$0")")"
ADDON_ZIP="$MY_DIR/web/static/flamenco-addon.zip"
WORKER_TARGET=/shared/software/flamenco3-worker/flamenco-worker

TIMESTAMP=$(date +'%Y-%m-%d-%H%M%S')

set -e

function prompt() {
  echo
  echo -------------------
  printf " \033[38;5;214m$@\033[0m\n"
  echo -------------------
  echo
}

prompt "Building Flamenco"
make

prompt "Deploying Manager"
ssh -o ClearAllForwardings=yes flamenco.farm.blender -t sudo systemctl stop flamenco3-manager
ssh -o ClearAllForwardings=yes flamenco.farm.blender -t cp /home/flamenco3/flamenco-manager.sqlite /home/flamenco3/flamenco-manager.sqlite-bak-$TIMESTAMP
scp flamenco-manager flamenco.farm.blender:/home/flamenco3/
ssh -o ClearAllForwardings=yes flamenco.farm.blender -t sudo systemctl start flamenco3-manager

prompt "Deploying Worker"
if [ -e "$WORKER_TARGET" ]; then
  # Move the old worker out of the way. This should keep the old executable
  # as-is on disk, hopefully keeping currently-running processes happy.
  mv "$WORKER_TARGET" "$WORKER_TARGET.old"
fi
cp -f flamenco-worker "$WORKER_TARGET"

prompt "Deploying Blender Add-on"
rm -rf /shared/software/addons/flamenco
pushd /shared/software/addons
unzip -q "$ADDON_ZIP"
popd

prompt "Done!"
echo "Deployment done, be sure to restart all the Workers and poke Artists to reload their Blender add-on."
echo

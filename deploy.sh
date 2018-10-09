#!/bin/sh

set -o errexit
set -o pipefail

echo "Installing: Goofys"
cp /goofys /target/.tmp_mount
mv -f /target/.tmp_mount /target/mount
chmod +x /target/mount

echo "Installing: Flexvolume"
cp /flexvolume /target/.tmp_goofys
mv -f /target/.tmp_goofys /target/goofys
chmod +x /target/goofys

while : ; do
  sleep 3600
done
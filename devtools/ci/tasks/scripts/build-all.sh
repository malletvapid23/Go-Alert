#!/bin/sh
set -ex
PREFIX=$1
make check test smoketest cy-wide-prod-run cy-mobile-prod-run bin/goalert BUNDLE=1 CI=1 DB_URL=$DB_URL
./bin/goalert self-test --offline
VERSION=$(./bin/goalert version | head -n 1 |awk '{print $2}')

for PLATFORM in darwin-amd64 linux-amd64 linux-arm linux-arm64
do
    SRC=bin/goalert-${PLATFORM}.tgz
    make $SRC
    cp $SRC ${PREFIX}goalert-${VERSION}-${PLATFORM}.tgz
done

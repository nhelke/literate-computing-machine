#!/usr/bin/env bash

set -e -u

pushd "$(dirname "$0")" > /dev/null
for delay in 10s 10m 1h 2h 3h 4h 5h 6h 7h 8h 9h; do
  time go run main.go "$delay" &
done
wait
popd > /dev/null
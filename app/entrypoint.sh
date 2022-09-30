#!/usr/bin/env bash
set -Eeuo pipefail

subcommand="${1:-}"

case "$subcommand" in
  hang)
    echo "HANG"
    tail -f /dev/null & wait
  ;;
esac

JETSTREAMER_SERVER_NAME=$HOSTNAME
export JETSTREAMER_SERVER_NAME

natsCmd="nats-server --name $JETSTREAMER_SERVER_NAME --config /app/nats-server.conf $*"
echo "entrypoint.sh starting $natsCmd"
exec $natsCmd

#!/bin/sh
set -eu

# Start the wg-easy entrypoint if present, otherwise try to run wg-easy binary.
# The entrypoint is started in background so we can also run the Go app side-by-side.
# This script forwards INT/TERM to both processes and waits for either to exit.

start_wg_easy() {
  echo "Starting wg-easy"
  /usr/bin/dumb-init node server.js "$@" &
  WG_PID=$!
}

start_app() {
  echo "Starting clearnetNode"
  /app/clearnetNodeJodAlpine &
  APP_PID=$!
}

# Start services
WG_PID=0
APP_PID=0
start_wg_easy "$@"
sleep 6
start_app

# trap signals and forward to children
_term() {
  echo "Terminating processes..."
  [ "$APP_PID" -ne 0 ] && kill -TERM "$APP_PID" 2>/dev/null || true
  [ "$WG_PID" -ne 0 ] && kill -TERM "$WG_PID" 2>/dev/null || true
}

trap _term INT TERM

# Wait for either process to exit; if one exits, forward and exit with its status.
wait_for_pids() {
  while true; do
    if ! kill -0 "$APP_PID" 2>/dev/null; then
      wait "$APP_PID" || rc=$?
      echo "App exited with ${rc:-0}, stopping wg-easy"
      [ "$WG_PID" -ne 0 ] && kill -TERM "$WG_PID" 2>/dev/null || true
      exit ${rc:-0}
    fi
    if ! kill -0 "$WG_PID" 2>/dev/null; then
      wait "$WG_PID" || rc=$?
      echo "wg-easy exited with ${rc:-0}, stopping app"
      [ "$APP_PID" -ne 0 ] && kill -TERM "$APP_PID" 2>/dev/null || true
      exit ${rc:-0}
    fi
    sleep 1
  done
}

wait_for_pids
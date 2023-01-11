#!/bin/sh

set -e

echo "run gooseup"
/app/migrate -dir ./migration/ -v postgres "$DB_SOURCE" up

echo "start the app"
exec "$@"
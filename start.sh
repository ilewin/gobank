#!/bin/sh
set -e
echo "Run DB Migration Now"
cat /app/app.env
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up
echo "Run Server"
exec "$@"

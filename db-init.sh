#!/bin/bash
set -e

echo "Waiting for postgres to start..."

# Run migration scripts (excluding down migrations)
for f in /tmp/psql_data/migrations/*up.sql; do
  echo "Running migration $f"
  psql -U postgres -d keyonthego -v ON_ERROR_STOP=1 -f "$f"
done

# Run seeding script
for f in /tmp/psql_data/seed/*.sql; do
  echo "Running seeding $f"
  psql -U postgres -d keyonthego -v ON_ERROR_STOP=1 -f "$f"
done

echo "Postgres is up - waiting for requests"
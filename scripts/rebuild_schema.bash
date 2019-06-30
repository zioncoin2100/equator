#! /usr/bin/env bash
set -e

# This scripts rebuilds the latest.sql file included in the schema package.

gb generate github.com/zion/equator/db2/schema
gb build
dropdb equator_schema --if-exists
createdb equator_schema
DATABASE_URL=postgres://localhost/equator_schema?sslmode=disable ./bin/equator db migrate up

DUMP_OPTS="--schema=public --no-owner --no-acl --inserts"
LATEST_PATH="src/github.com/zion/equator/db2/schema/latest.sql"
BLANK_PATH="src/github.com/zion/equator/test/scenarios/blank-equator.sql"

pg_dump postgres://localhost/equator_schema?sslmode=disable $DUMP_OPTS \
  | sed '/SET idle_in_transaction_session_timeout/d'  \
  | sed '/SET row_security/d' \
  > $LATEST_PATH
pg_dump postgres://localhost/equator_schema?sslmode=disable \
  --clean --if-exists $DUMP_OPTS \
  | sed '/SET idle_in_transaction_session_timeout/d'  \
  | sed '/SET row_security/d' \
  > $BLANK_PATH

gb generate github.com/zion/equator/db2/schema
gb generate github.com/zion/equator/test
gb build

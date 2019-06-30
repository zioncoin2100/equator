#! /usr/bin/env bash
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PACKAGES=$(find src/github.com/zion/equator/test/scenarios -iname '*.rb' -not -name '_common_accounts.rb')
# PACKAGES=$(find src/github.com/zion/equator/test/scenarios -iname 'kahuna.rb')

gb build

dropdb hayashi_scenarios --if-exists
createdb hayashi_scenarios

export ZION_CORE_DATABASE_URL="postgres://localhost/hayashi_scenarios?sslmode=disable"
export DATABASE_URL="postgres://localhost/equator_scenarios?sslmode=disable"
export NETWORK_PASSPHRASE="Test SDF Network ; September 2015"
export ZION_CORE_URL="http://localhost:8080"
export SKIP_CURSOR_UPDATE="true"

# run all scenarios
for i in $PACKAGES; do
  CORE_SQL="${i%.rb}-core.sql"
  HORIZON_SQL="${i%.rb}-equator.sql"
  bundle exec scc -r $i --dump-root-db > $CORE_SQL

  # load the core scenario
  psql $ZION_CORE_DATABASE_URL < $CORE_SQL

  # recreate equator dbs
  dropdb equator_scenarios --if-exists
  createdb equator_scenarios

  # import the core data into equator
  $DIR/../bin/equator db init
  $DIR/../bin/equator db reingest

  # write equator data to sql file
  pg_dump $DATABASE_URL \
    --clean --if-exists --no-owner --no-acl --inserts \
    | sed '/SET idle_in_transaction_session_timeout/d' \
    | sed '/SET row_security/d' \
    > $HORIZON_SQL
done


# commit new sql files to bindata
gb generate github.com/zion/equator/test/scenarios
# gb test github.com/zion/equator/ingest

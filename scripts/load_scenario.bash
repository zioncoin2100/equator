#! /usr/bin/env bash

set -e

SCENARIO=$1
CORE_SQL=./src/github.com/zion/equator/test/scenarios/$SCENARIO-core.sql
HORIZON_SQL=./src/github.com/zion/equator/test/scenarios/$SCENARIO-equator.sql

echo "psql $ZION_CORE_DATABASE_URL < $CORE_SQL" 
psql $ZION_CORE_DATABASE_URL < $CORE_SQL 
echo "psql $DATABASE_URL < $HORIZON_SQL"
psql $DATABASE_URL < $HORIZON_SQL 

#!/bin/bash

# Script that runs when Postgres container spins up

# set -e

# psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
#	GRANT ALL PRIVILEGES ON DATABASE baetyl TO pascalallen;
# EOSQL

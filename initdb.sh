#!/bin/bash

echo "executing initdb.sh..."

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE TABLE trademarks(
        ID SERIAL           PRIMARY KEY,
        NAME                TEXT    NOT NULL,
        STATUS_CODE         TEXT    NOT NULL,
        STATUS_DATE         DATE    NOT NULL
    );
    COPY trademarks(name, status_code, status_date)
    FROM '/docker-entrypoint-initdb.d/output.csv'
    DELIMITER ','
    CSV HEADER;
    CREATE EXTENSION fuzzystrmatch;
    CREATE INDEX idx_trademarks_name ON trademarks(name);
EOSQL

echo "database initialized"
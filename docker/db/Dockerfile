FROM postgres:15 as vldemo-db

RUN mkdir -p /docker-entrypoint-initdb.d

COPY initdb.sh /docker-entrypoint-initdb.d

COPY temp/output.csv /docker-entrypoint-initdb.d

COPY docker/db/pg_pass /run/secrets/postgres-passwd
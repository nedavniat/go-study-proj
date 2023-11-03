#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  DROP TABLE IF EXISTS tasks;

  CREATE TABLE tasks (
    id SERIAL,
    title character varying(255) NOT NULL,
      content character varying(255) NOT NULL,
      due_date timestamp NOT NULL,
      CONSTRAINT PK_TASKS PRIMARY KEY (id)
  );
EOSQL

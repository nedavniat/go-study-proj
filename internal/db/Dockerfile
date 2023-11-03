FROM postgres:12-alpine

RUN mkdir -p /docker-entrypoint-initdb.d
COPY ./initdb.sh /docker-entrypoint-initdb.d/


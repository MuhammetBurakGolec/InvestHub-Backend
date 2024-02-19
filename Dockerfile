FROM postgres:latest

COPY ./db/init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432

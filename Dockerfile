FROM mysql:latest

ADD create-tables.sql /docker-entrypoint-initdb.d

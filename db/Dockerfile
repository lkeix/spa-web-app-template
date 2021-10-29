FROM postgres:12-alpine
MAINTAINER lkeix

ENV POSTGRES_USER=root
ENV POSTGRES_PASSWORD=root
ENV POSTGRES_DB=sample_db
ENV POSTGRES_INITDB_ARGS=--encoding=UTF-8


ADD ./db_schema.sql /docker-entrypoint-initdb.d

EXPOSE 5432
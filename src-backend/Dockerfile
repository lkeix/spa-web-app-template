FROM golang:alpine
MAINTAINER lkeix

ENV DB_TYPE postgres
ENV DB_HOST postgres
ENV DB_PORT 5432
ENV DB_USER root
ENV DB_PASSWORD root
ENV DB_NAME sample_db
ENV DB_SCHEMA sample_schema
ENV DB_SSL disable

WORKDIR /go/src/app

COPY . /go/src/app
# RUN go build -o serve serve.go

# CMD ["./serve"]

RUN go install
RUN go get -u github.com/cosmtrek/air

CMD ["air"]
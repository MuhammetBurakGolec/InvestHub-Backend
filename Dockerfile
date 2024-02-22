FROM golang:1.22-bullseye

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/bin/app

EXPOSE 5001

CMD ["/go/bin/app"]


# Veri tabanı İşlemleri Aktarılacak
# FROM postgres:latest
# COPY ./db/init.sql /docker-entrypoint-initdb.d/
# EXPOSE 5432

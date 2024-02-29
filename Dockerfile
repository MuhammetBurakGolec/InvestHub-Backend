FROM golang:1.22-bullseye

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main .

EXPOSE 5001

CMD ["./main"]


# Veri tabanı İşlemleri Aktarılacak
# FROM postgres:latest
# COPY ./db/init.sql /docker-entrypoint-initdb.d/
# EXPOSE 5432

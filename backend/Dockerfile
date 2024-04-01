FROM golang:latest

WORKDIR /app

COPY ./go-api /app

RUN echo "DB_HOST=postgres-db\nDB_PORT=5432\nDB_USER=postgres\nDB_PASSWORD=password\nDB_NAME=nbaappdb\nBACKEND_PORT=9000" > /app/.env

RUN go build -o main .

EXPOSE 9000

CMD ["/app/main"]
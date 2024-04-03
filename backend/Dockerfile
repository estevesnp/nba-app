FROM golang:latest

WORKDIR /app

COPY ./go-api /app

COPY ./docker.env /app/.env

RUN go build -o main .

EXPOSE 9000

CMD ["/app/main"]
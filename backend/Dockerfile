FROM golang:latest

WORKDIR /app

COPY ./go-api /app

RUN go build -o main .

EXPOSE 9000

CMD ["/app/main"]
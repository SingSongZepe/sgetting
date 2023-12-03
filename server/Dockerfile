FROM golang:1.16

WORKDIR /app

COPY /app/server.go .

RUN go build -o server server.go

EXPOSE 8080

CMD ["./server"]
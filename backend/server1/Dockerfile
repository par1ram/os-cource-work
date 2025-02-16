FROM golang:1.23 as builder
WORKDIR /app

COPY . .

RUN cd server1 && go mod tidy && go build -o server1 cmd/main.go

FROM debian:latest
WORKDIR /app

COPY --from=builder /app/server1/server1 .

CMD ["./server1"]

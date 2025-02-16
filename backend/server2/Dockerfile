FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN cd server2 && go mod tidy && go build -o server2 cmd/main.go

FROM debian:latest
WORKDIR /app
COPY --from=builder /app/server2/server2 .

CMD ["./server2"]

FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN cd client && go mod tidy && go build -o client cmd/main.go

FROM debian:latest
WORKDIR /app
COPY --from=builder /app/client/client .

CMD ["./client"]

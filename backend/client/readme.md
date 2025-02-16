go build -o client && ./client

docker-compose build
docker-compose run client

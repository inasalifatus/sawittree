.PHONY: init test test-api build run docker-up docker-down

init:
	go mod tidy
	go generate ./...

test:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out | grep total | awk '{print "Coverage:", $$3}'

test-api:
	# Tambahkan perintah untuk menjalankan API tests
	echo "Running API tests..."

build:
	go build -o app .

run:
	./app

docker-up:
	docker compose up --build -d
	sleep 30
	make test-api

docker-down:
	docker compose down --volumes

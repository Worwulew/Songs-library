# Main

start: migrate_up run

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -v -cover ./...

#Migrator

migrate_up:
	go run ./cmd/migrator/main.go up

migrate_down:
	go run ./cmd/migrator/main.go down

#Swagger

swag:
	swag init -g cmd/api/main.go

#Docker

docker:
	docker-compose -f docker-compose.yml up --build
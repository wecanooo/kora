APP_NAME = "kora"

default:
	go build -o ${APP_NAME}

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

redis:
	docker run --name redis -p 6379:6379 -d redis:alpine

install:
	go mod download

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root kora_development

dropdb:
	docker exec -it postgres12 dropdb kora_development

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir db/migrate $$name

migrateup:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/kora_development?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/kora_development?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/kora_development?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/kora_development?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go server

help:
	@echo "make - compile the source code"
	@echo "make postgres - launch postgres docker container"
	@echo "make redis - launch redis docker container"
	@echo "make install - install dep"
	@echo "make createdb - create a database"
	@echo "make dropdb - drop a database"
	@echo "make migration - create migration file"
	@echo "make migrateup - migrate up all migration files"
	@echo "make migrateup1 - migrate up recent one migration file"
	@echo "make migratedown - migrate down all migration files"
	@echo "make migratedown1 - migrate down recent one migration files"
	@echo "make sqlc - generate type safe go codes by sqlc generate"
	@echo "make server - start a server"

.PHONY: postgres redis install createdb dropdb migration migrateup migrateup1 migratedown migratedown1 sqlc test server help
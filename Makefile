postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

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

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/wecanooo/mars/model Store

.PHONY: postgres createdb dropdb migration migrateup migrateup1 migratedown migratedown1 sqlc test server mock
database:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root kora

dropdb:
	docker exec -it postgres12 dropdb kora

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir db/migrate $$name

migrateup:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/kora?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/kora?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/wecanooo/mars/model Store

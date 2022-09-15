postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret	-d postgres 

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb --username=root simple_bank

migrateup1:
	migrate -path "db/migration" -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migrateup:
	migrate -path "db/migration" -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown1:
	migrate -path "db/migration" -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown:
	migrate -path "db/migration" -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/transparentideas/gobank/db/sqlc Store

.PHONY:postgres createdb dropdb	migrateup migratedown migrateup1 migratedown1 sqlc server mock
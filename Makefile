postgres:
	docker run --name postgres --network assignment-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -dp 8081:5432 postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root assignment

dropdb:
	docker exec -it dropdb assignment

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/assignment?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/assignment?sslmode=disable" -verbose down

server:
	go run main.go

dev:
	nodemon

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown server dev sqlc test
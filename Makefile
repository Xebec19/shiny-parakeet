postgres:
	sudo docker run --name admybrand -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 8081:5432 -d postgres:12-alpine

createdb:
	sudo docker exec -it admybrand createdb --username=root --owner=root assignment

dropdb:
	sudo docker exec -it dropdb assignment

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:8081/assignment?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:8081/assignment?sslmode=disable" -verbose down

server:
	go run main.go

dev:
	nodemon

.PHONY: postgres createdb dropdb migrateup migratedown server dev
postgres:
	sudo docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:17-alpine

createdb:
	sudo docker exec -it postgres17 createdb --username=root --owner=root user_management

dropdb:
	sudo docker exec -it postgres17 dropdb user_management

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/user_management?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/user_management?sslmode=disable" -verbose down



.PHONY: createdb

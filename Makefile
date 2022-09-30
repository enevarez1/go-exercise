postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

pgadmin:
	docker run --name pgadmin -p 80:80 -e 'PGADMIN_DEFAULT_EMAIL=root@test.com' -e 'PGADMIN_DEFAULT_PASSWORD=secret' -d dpage/pgadmin4

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root exercise_db

dropdb:
	docker exec -t postgres13 dropdb exercise_db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/exercise_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/exercise_db?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

.PHONY: postgres pgadmin createdb dropdb migrateup migratedown test sqlc
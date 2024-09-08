DB_URL=postgresql://root:root@localhost:5432/golang_project?sslmode=disable

postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15-alpine

start:
	docker start postgres15 pgadmin
stop:
	docker stop postgres15 pgadmin

pgadmin:
	docker run --name=pgadmin -p 5050:80 -e PGADMIN_DEFAULT_EMAIL=a@a.com -e PGADMIN_DEFAULT_PASSWORD=root -d dpage/pgadmin4

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root golang_project

dropdb:
	docker exec -it postgres15 dropdb golang_project

migrateup:
	migrate -path ./db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path ./db/migrations -database "$(DB_URL)" -verbose down

migratedrop:
	migrate -path ./db/migrations -database "$(DB_URL)" -verbose drop -f

# seed:
#     docker cp ./db/seed/dummy_data.sql db:/tmp/query_file.sql | docker exec -T db psql "${DB_URL}" -f /tmp/query_file.sql	
sqlc:
	sqlc generate

.PHONY: postgres start stop pgadmin createdb dropdb migrateup migratedown migratedrop sqlc seed
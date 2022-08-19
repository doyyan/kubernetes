DB_URL=postgresql://root:secret@localhost:5432/kubernetes?sslmode=disable

network:
	docker network create k8s-network

postgres:
	docker network create k8s-network
	docker run --name postgres --network k8s-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

rm-postgres:
	docker kill postgres; docker rm postgres; docker network rm k8s-network

createdb:
	docker exec -it postgres createdb kubernetes -O root
dropdb:
	docker exec -it postgres dropdb kubernetes

migrateup:
	migrate -path internal/app/adapter/postgresql/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path internal/app/adapter/postgresql/migration -database "$(DB_URL)" -verbose down

migrateup1:
	migrate -path internal/app/adapter/postgresql/migration -database "$(DB_URL)" -verbose up 1

migratedown1:
	migrate -path internal/app/adapter/postgresql/migration -database "$(DB_URL)" -verbose down 1

dockerbuild:
	go mod tidy -v
	go mod vendor
	docker build . -t kubectldb -f build/Dockerfile

test:
	go test -v -cover ./...

server:
	go run cmd/app/main.go

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 test server

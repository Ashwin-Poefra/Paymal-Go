migrate-up:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/paymal?sslmode=disable" -verbose up

migrate-up-1:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/paymal?sslmode=disable" -verbose up 1

migrate-down:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/paymal?sslmode=disable" -verbose down

migrate-down-1:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/paymal?sslmode=disable" -verbose down 1

test:
	go test -v -cover ./...

sqlc:
	docker run --rm -v $(CURDIR):/src -w /src kjconroy/sqlc generate

dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
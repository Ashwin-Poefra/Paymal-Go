test:
	go test -v -cover ./...

sqlc:
	docker run --rm -v $(CURDIR):/src -w /src kjconroy/sqlc generate

dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
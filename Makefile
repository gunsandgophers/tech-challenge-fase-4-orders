build:
	docker compose up -d --build

start:
	docker compose up -d

stop:
	docker compose down

logs/app:
	docker compose logs -f --no-log-prefix app

migrate:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path=./migrations/ -database ${DB_URL} up

migrate/create:
	docker run -v ./migrations:/migrations --network host migrate/migrate create -ext sql -dir ./migrations $(name)

swagger:
	docker run --rm -v ./:/code ghcr.io/swaggo/swag:latest init

swagger-mac:
	docker run --platform linux/amd64 --rm -v ./:/code ghcr.io/swaggo/swag:latest init

mock/create:
	docker run -v ./:/src -w /src vektra/mockery --all

test:
	go test ./internal/core/...

test/coverage:
	go test ./internal/core/... -coverprofile=cover.out

test/output:
	go tool cover -html=cover.out

test/bdd:
	go test -v ./internal/tests/bdd/...

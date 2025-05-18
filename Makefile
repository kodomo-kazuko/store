run:
	@air -c ./.air.toml

query:
	go run database/query/main.go

migrate:
	go run database/migrate/main.go

local:
	go build main.go && main -ip 10.147.17.168:8080

seed:
	go run database/seed/main.go

all: migrate
	make query
	make seed
	make local

diff:
	atlas migrate diff --env gorm

apply:
	atlas migrate apply --env gorm

.PHONY: pgr postgrest docker

pgr: postgrest docker

postgrest:
	@echo "Starting PostgREST..."
	postgrest start.conf

docker:
	@echo "Starting Docker Compose..."
	docker-compose up
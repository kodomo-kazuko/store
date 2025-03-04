run:
	@air -c ./.air.toml

query:
	go run database/query/main.go

migrate:
	go run database/migrate/main.go

local:
	go build main.go && ./main -ip 10.147.17.117:8080

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

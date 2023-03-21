#!make
include .env

goose-build = ./goose -dir ./db/migrations postgres ${DATABASE_DNS}

all: dev

dev:
	nodemon --exec go run main.go --signal SIGTERM
	
db.status:
	$(goose-build) status

migration.create:
	$(goose-build) create "$(name)" go

migration.run:
	$(goose-build) up

migration.revert:
	$(goose-build) down

goose.build:
	go build -o goose ./cmd/goose
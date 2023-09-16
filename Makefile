up: pg

pg:
	docker run --name pg16 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=go_api -d postgres:16-alpine

gen:
	go generate

migration:
	@migrate create -ext sql -dir migrations -seq $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	go run main.go db migrate up

migrate-down:
	go run main.go db migrate down
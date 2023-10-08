up:
	docker compose up -d

fmt:
	go fmt ./...
	
gen:
	go generate

runserver:
	go run main.go server

migration:
	@migrate create -ext sql -dir migrations -seq $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	go run main.go db migrate up

migrate-down:
	go run main.go db migrate down
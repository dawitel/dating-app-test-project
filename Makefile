run:
	@go run ./cmd/main.go

build:
	@docker build -t matchmaking-app .

migrate-up:
	@migrate -path ./migrations -database "postgresql://root:1234@localhost:5432/dating_app?sslmode=disable" -verbose up

migrate-down:
	@migrate -path ./migrations -database "postgresql://root:1234@localhost:5432/dating_app?sslmode=disable" down

migrate-create:
	@ @migrate create -ext sql -dir ./migrations -seq init_schema
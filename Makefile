# Define environment variables
DB_URL ?= "postgresql://root:1234@localhost:5432/dating_app?sslmode=disable"
IMAGE_NAME ?= "matchmaking-app"
PORT ?= "8080"

# Run the Go application
run:
	@go run ./cmd/main.go

# Build the Docker image
build:
	@docker build -t $(IMAGE_NAME) .

# Apply database migrations
migrate-up:
	@migrate -path ./migrations -database "$(DB_URL)" -verbose up

migrate-down:
	@migrate -path ./migrations -database "$(DB_URL)" down

migrate-create:
	@migrate create -ext sql -dir ./migrations -seq init_schema

# Start air for live reloading
air:
	@air

# Docker Compose commands
docker-compose-up:
	@docker-compose up --build

docker-compose-down:
	@docker-compose down

docker-compose-build:
	@docker-compose build --no-cache

# Start the application and run migrations
start: docker-compose-up migrate-up

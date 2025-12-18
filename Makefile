-include .env
export
BINARY=go-social-feed

run-api:
	@echo "Running the application"
	go run cmd/api/main.go

swagger:
	@echo "Generating swagger documentation..."
	swag init -g cmd/api/main.go -o docs

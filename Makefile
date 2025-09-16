include .env
export
BINARY=go-social-feed

## Run the application
run-api:
	@echo "Running the application"
	@make swagger
	@go run cmd/api/main.go

## Generate swagger documentation
swagger:
	@echo "Generating swagger documentation..."
	@swag init -g cmd/api/main.go -o docs
	@echo "Swagger docs generated successfully"
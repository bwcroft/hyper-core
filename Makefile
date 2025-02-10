.PHONY: help api	

help: 
	@echo "Available commands:"
	@echo ""
	@echo "api  - Start development environment"
	@echo "test - Run all tests"

api:
	go run cmd/api/main.go -env=./.env

test:
	go test ./...

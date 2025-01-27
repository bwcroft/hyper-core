.PHONY: help api	

help: 
	@echo "Available commands:"
	@echo ""
	@echo "api - Start development environment"

api:
	go run cmd/api/main.go -env=./.env

.PHONY: help api	

help: 
	@echo "Available commands:"
	@echo ""
	@echo "test    Run all tests"
	@echo "test-v  Run all tests verbose"

test:
	go test ./...

test-v:
	go test ./... -v

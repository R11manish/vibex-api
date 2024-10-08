# Define variables
APP_NAME = vibex-api
SRC_DIR = .
BUILD_DIR = build
GO_FILES = $(wildcard $(SRC_DIR)/*.go)


.PHONY: run


run: 
	go run ${SRC_DIR}/cmd/api/main.go



# Help target
.PHONY: help
help:
	@echo "Makefile for $(APP_NAME)"
	@echo "Available targets:"
	@echo "  all     - Build and run the application"
	@echo "  build   - Build the application"
	@echo "  run     - Run the application"
	@echo "  clean   - Remove generated files"
	@echo "  help    - Show this help message"
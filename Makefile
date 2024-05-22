# Makefile

APP_NAME=hka-server-login

# Detect the operating system
ifeq ($(OS),Windows_NT)
    # Windows settings
    APP_DIR=.\app
    BUILD_DIR=.\bin
    MKDIR = if not exist $(BUILD_DIR) mkdir $(BUILD_DIR)
    RMDIR = if exist $(BUILD_DIR) rmdir /s /q $(BUILD_DIR)
    EXE = .exe
    SEP = \\
else
    # Unix settings
    APP_DIR=./app
    BUILD_DIR=./bin
    MKDIR = mkdir -p $(BUILD_DIR)
    RMDIR = rm -rf $(BUILD_DIR)
    EXE =
    SEP = /
endif

.PHONY: all build run clean docker-build docker-run docker-clean

all: build

build:
	@echo "Building application..."
	@$(MKDIR)
	@cd $(APP_DIR) && go build -o ..$(SEP)$(BUILD_DIR)$(SEP)$(APP_NAME)$(EXE) ./cmd/main

run: build
	@echo "Running application..."
	@$(BUILD_DIR)$(SEP)$(APP_NAME)$(EXE) --configPath "./app/config/config.json"

clean:
	@echo "Cleaning up..."
	@go clean
	@$(RMDIR)

test:
	@echo "Running tests..."
	@cd $(APP_DIR) && go test ./...

docker-build:
	@echo "Building the Docker image..."
	@docker build -t $(APP_NAME) .

docker-run:
	@echo "Running the Docker container..."
	@docker run --rm -p 8080:8080 $(APP_NAME)

docker-clean:
	@echo "Cleaning up Docker images..."
	@docker rmi $(APP_NAME)
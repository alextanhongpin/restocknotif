include .env
export # Export the environment variables.

# Include all other sub-makefiles.
include *.mk


# Spins up all containers.
up:
	@docker-compose up -d


# Spins down all containers.
down:
	@docker-compose down


# Starts the application.
start:
	@go run main.go

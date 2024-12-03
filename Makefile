# Variables
COMPOSE_FILE=docker-compose.yml

# Build all services
build:
	docker-compose -f $(COMPOSE_FILE) build

# Start all services
up:
	docker-compose -f $(COMPOSE_FILE) up -d

# Stop all services
down:
	docker-compose -f $(COMPOSE_FILE) down

# Restart all services
restart: down up

# View logs
logs:
	docker-compose -f $(COMPOSE_FILE) logs -f

# Clean up all containers, networks, and volumes
clean:
	docker-compose -f $(COMPOSE_FILE) down -v

# Check the status of all services
ps:
	docker-compose -f $(COMPOSE_FILE) ps

# Execute command in auth-service container
exec-auth:
	docker-compose -f $(COMPOSE_FILE) exec auth-service /bin/sh

# Execute command in blog-service container
exec-blog:
	docker-compose -f $(COMPOSE_FILE) exec blog-service /bin/sh

# Run tests
test:
	# Implement commands to test your services, e.g., docker-compose -f $(COMPOSE_FILE) run test

# Default goal
.PHONY: default
default: up
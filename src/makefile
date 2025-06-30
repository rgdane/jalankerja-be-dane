.PHONY: help up down restart build rebuild logs clean ps shell

# Default target
help:
	@echo "Available commands:"
	@echo "  make run         - Start all services"
	@echo "  make stop        - Stop all services"
	@echo "  make restart     - Restart all services (down + up)"
	@echo "  make build       - Build services"
	@echo "  make rebuild     - Rebuild services without cache"
	@echo "  make logs        - Show logs"
	@echo "  make ps          - Show running containers"
	@echo "  make clean       - Clean up containers, images, and volumes"
	@echo "  make shell       - Enter nestjs-app container shell"

# Start services
run:
	docker-compose up -d

# Stop services  
stop:
	docker-compose down

# Restart services
restart: down up

# Build services
build:
	docker-compose build

# Rebuild without cache
rebuild:
	docker-compose build --no-cache

# Show logs
logs:
	docker-compose logs -f

# Show running containers
ps:
	docker-compose ps

# Clean up everything
clean:
	docker-compose down -v --remove-orphans
	docker system prune -f

# Enter container shell
shell:
	docker exec -it src-nestjs-app-1 sh

# Quick restart with rebuild
fresh: down rebuild up
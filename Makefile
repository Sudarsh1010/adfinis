.PHONY: help dev dev-backend dev-frontend clean build build-frontend build-server

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

dev: ## Start both backend and frontend with hot reload
	@echo "🚀 Starting development environment..."
	@echo "   Backend:  http://localhost:3000 (Go + Air)"
	@echo "   Frontend: http://localhost:5173 (SvelteKit + Vite)"
	@echo ""
	@./dev.sh

dev-backend: ## Start backend only with hot reload
	@echo "🔧 Starting backend with Air..."
	@air

dev-frontend: ## Start frontend only with dev server
	@echo "🎨 Starting frontend dev server..."
	@cd web && pnpm run dev

clean: ## Clean temporary files
	@echo "🧹 Cleaning temporary files..."
	@rm -rf ./tmp/*
	@echo "✓ Cleaned tmp directory"

build-frontend: ## Build frontend for production
	@echo "📦 Building frontend..."
	@cd web && pnpm install && pnpm run build
	@echo "✓ Frontend built in web/build/"

build-server: build-frontend ## Build server binary with embedded frontend
	@echo "🔨 Building server binary..."
	@go build -o bin/server ./cmd/server
	@echo "✓ Server binary: bin/server"

build: build-server ## Full production build

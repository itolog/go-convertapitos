# ConvertApiTos

A web application for converting API specifications with a Vue.js frontend and Go backend.

## Requirements

- **Go** 1.24+ 
- **Node.js** 22+ with npm
- **Docker** and **Docker Compose**
- **PostgreSQL** (for local development without Docker)

## Quick Start

### 1. Clone the repository

```bash
git clone <repository-url>
cd go-convertapitos
````

```shell
cp .env.example .env
```

## Local Development
### Option 1: With Docker (recommended)

# Start development environment with hot reload
```bash
make docker-up-dev
```
# Stop containers
```bash
make docker-stop-dev
```
# Remove containers completely
```bash
make docker-down-dev
```
After startup:
- **API**: [http://localhost:3000](http://localhost:3000)
- **Frontend**: [http://localhost:5173](http://localhost:5173)
- **PostgreSQL**: localhost:5432

### Option 2: Local setup

# Install dependencies
```bash
go mod download
```
# Run migrations (requires PostgreSQL)
```bash
make migrations-auto
```
# Start with hot reload
```bash
make watch
```
# Or regular start
make

#### Frontend
cd frontend
rm -rf node_modules package-lock.json
npm install

# Start dev server
```bash
npm run dev
```
## Production Mode
### Docker (recommended)

# Build frontend and start production containers
```bash
make docker-up-with-frontend
```
# Or start without rebuilding frontend
```bash
make docker-up
```
# Apply migrations in production
```bash
make docker-migrations-prod
```
# Stop
```bash
make docker-stop
```
# Complete cleanup
```bash
make docker-down
```
### Manual build

# Build frontend
```bash
make frontend-build
```
# Build backend
```bash
make build
```
# Apply migrations
```bash
make migrations-auto-prod
```
# Start production server
```bash
make run-prod
```
## Useful Commands
# Hot reload with air
```bash
make watch
```
or
```bash
make
```
# Run with garbage collector flags
```bash
make run-gcf
```
# Frontend linting
cd frontend && npm run lint

# Frontend code formatting
cd frontend && npm run format

### Development
# Hot reload with air
```bash
make watch
```
# Run with garbage collector flags
```bash
make run-gcf
```
# Frontend linting
cd frontend && npm run lint

# Frontend code formatting
cd frontend && npm run format

### Database
# Migrations for development
```bash
make migrations-auto
```
# Migrations for production
```bash
make migrations-auto-prod
```
# Migrations in Docker dev
```bash
make docker-migrations-dev
```
# Migrations in Docker prod
```bash
make docker-migrations-prod
```
### Swagger documentation
# Initialize Swagger
```bash
make swagger-init
```
# Format Swagger
```bash
make swagger-fmt
```
### Frontend
cd frontend
rm -rf node_modules package-lock.json
npm install

# Generate API types from Swagger
```bash
npm run gen-api-types
```
# Check dependency updates
```bash
npm run check-deps
```
# Update dependencies
```bash
npm run up-deps
```
# Build for production
```bash
npm run build
```
# Preview build
```bash
npm run preview
```

## API Documentation
After starting the application, Swagger documentation is available at:
- **Development**: [http://localhost:3000/api/v1/swagger/](http://localhost:3000/api/v1/swagger/)
- **Production**: [http://your-domain/api/v1/swagger/](http://your-domain/api/v1/swagger/)

## Environment Variables
Main environment variables (see ): `.env.example`

## Troubleshooting
### Database issues
# Recreate containers with data cleanup
make docker-down-dev
docker volume prune
make docker-up-dev

### Frontend issues
cd frontend
rm -rf node_modules package-lock.json
npm install

### Go modules issues
go clean -modcache
go mod download

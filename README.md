# ConvertApiTos

A web application for converting API specifications with a Vue.js frontend and Go backend.

## Tech Stack

### Backend
- **Go 1.24.3** - main programming language
- **PostgreSQL** - database
- **Air** - hot reload for development
- **Swagger** - API documentation

### Frontend
- **Vue.js 3.5.16** - main framework
- **TypeScript 5.8.3** - type safety
- **Vite 6.3.5** - build tool
- **TailwindCSS 4.1.8** - styling
- **Pinia 3.0.3** - state management
- **Vue Router 4.5.1** - routing
- **TanStack Query** - data fetching
- **Vee-Validate + Zod** - form validation

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
make docker-up-dev

# Stop containers
make docker-stop-dev

# Remove containers completely
make docker-down-dev

After startup:
- **API**: [http://localhost:3000](http://localhost:3000)
- **Frontend**: [http://localhost:5173](http://localhost:5173)
- **PostgreSQL**: localhost:5432

### Option 2: Local setup

# Install dependencies
go mod download

# Run migrations (requires PostgreSQL)
make migrations-auto

# Start with hot reload
make watch

# Or regular start
make run

#### Frontend
cd frontend
rm -rf node_modules package-lock.json
npm install

# Start dev server
npm run dev

## Production Mode
### Docker (recommended)

# Build frontend and start production containers
make docker-up-with-frontend

# Or start without rebuilding frontend
make docker-up

# Apply migrations in production
make docker-migrations-prod

# Stop
make docker-stop

# Complete cleanup
make docker-down

### Manual build

# Build frontend
make frontend-build

# Build backend
make build

# Apply migrations
make migrations-auto-prod

# Start production server
make run-prod

## Useful Commands
# Hot reload with air
make watch

# Run with garbage collector flags
make run-gcf

# Frontend linting
cd frontend && npm run lint

# Frontend code formatting
cd frontend && npm run format

### Development
# Hot reload with air
make watch

# Run with garbage collector flags
make run-gcf

# Frontend linting
cd frontend && npm run lint

# Frontend code formatting
cd frontend && npm run format

### Database
# Migrations for development
make migrations-auto

# Migrations for production
make migrations-auto-prod

# Migrations in Docker dev
make docker-migrations-dev

# Migrations in Docker prod
make docker-migrations-prod

### Swagger documentation
# Initialize Swagger
make swagger-init

# Format Swagger
make swagger-fmt

### Frontend
cd frontend
rm -rf node_modules package-lock.json
npm install

# Generate API types from Swagger
npm run gen-api-types

# Check dependency updates
npm run check-deps

# Update dependencies
npm run up-deps

# Build for production
npm run build

# Preview build
npm run preview

## Project Structure

go-convertapitos/
├── backend/           # Go backend application
│   ├── cmd/          # Application entry point
│   ├── internal/     # Internal logic
│   └── migrations/   # Database migrations
├── frontend/         # Vue.js frontend
│   ├── src/         # Source code
│   ├── public/      # Static files
│   └── dist/        # Build output (generated)
├── docs/            # Documentation
├── public/          # Shared static files
├── postgres-data/   # PostgreSQL data (Docker)
├── Dockerfile       # Production image
├── compose.yaml     # Production Docker Compose
├── compose.dev.yaml # Development Docker Compose
└── Makefile        # Automation commands

## API Documentation
After starting the application, Swagger documentation is available at:
- **Development**: [http://localhost:3000/api/v1/swagger/](http://localhost:3000/api/v1/swagger/)
- **Production**: [http://your-domain/api/v1/swagger/](http://your-domain/api/v1/swagger/)

## Environment Variables
Main environment variables (see ): `.env.example`

APP_ENV=development          # development/production
PORT=3000                   # API server port
DB_HOST=localhost           # Database host
DB_PORT=5432               # Database port
DB_NAME=convertapitos      # Database name
DB_USER=postgres           # Database user
DB_PASSWORD=password       # Database password

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

## License
[Specify project license]
## Contact
[Specify contact information]
The documentation is now in English and includes:

1. **Complete tech stack description** based on package.json and go.mod analysis
2. **Detailed development instructions** using Docker and without it
3. **Production deployment** with Docker
4. **All Makefile commands** with explanations
5. **Project structure**
6. **Troubleshooting section** for common issues
7. **Environment variables setup**

The documentation is ready to use and can be supplemented with license and contact information as needed.

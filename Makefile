BINARY=convertapitos
CMD_DIR=./src/cmd
SWAGGER_FILE ?=./src/internal/router/router.go
COMOSE_DEV_FILE=compose.dev.yaml

# DEVELOPMENT
.PHONY: run
run:
	APP_ENV="development" go run ${CMD_DIR}

.PHONY: run-gcf
run-gcf:
	go run -gcflags '-m -l' ${CMD_DIR}

.PHONY: watch
watch:
	APP_ENV="development" air
# FRONTEND
.PHONY: frontend-build
frontend-build:
	cd frontend && npm install && npm run build

# BUILD
.PHONY: build
build:
	go build -o ./bin/${BINARY} ${CMD_DIR}

.PHONY: run-prod
run-prod:
	APP_ENV="production" ./bin/${BINARY}

# DOCKER DEV
.PHONY: docker-up-dev
docker-up-dev:
	docker compose -f ${COMOSE_DEV_FILE} up -d --build

.PHONY: docker-stop-dev
docker-stop-dev:
	docker compose -f ${COMOSE_DEV_FILE} stop

.PHONY: docker-down-dev
docker-down-dev:
	docker compose -f ${COMOSE_DEV_FILE} down

# DOCKER PROD
.PHONY: docker-up-with-frontend
docker-up-with-frontend: frontend-build
	docker compose up -d --build

.PHONY: docker-up
docker-up:
	docker compose up -d --build

.PHONY: docker-stop
docker-stop:
	docker compose stop

.PHONY: docker-down
docker-down:
	docker compose down

# DATABASE
.PHONY: migrations-auto
migrations-auto:
	APP_ENV="development" go run ./src/migrations

.PHONY: migrations-auto
migrations-auto-prod:
	APP_ENV="production" go run ./src/migrations

# SWAGGER
.PHONY: swagger-init
swagger-init:
	swag init -g ${SWAGGER_FILE}

.PHONY: swagger-fmt
swagger-fmt:
	swag fmt -g ${SWAGGER_FILE}


.DEFAULT_GOAL := watch
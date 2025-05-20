BINARY=convertapitos
CMD_DIR=./src/cmd
SWAGGER_FILE ?=./src/internal/router/router.go

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
# BUILD
.PHONY: build
build:
	go build -o ./bin/${BINARY} ${CMD_DIR}

.PHONY: run-prod
run-prod:
	APP_ENV="production" ./bin/${BINARY}

# DOCKER
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

# SWAGGER
.PHONY: swagger-init
swagger-init:
	swag init -g ${SWAGGER_FILE}

.PHONY: swagger-fmt
swagger-fmt:
	swag fmt -g ${SWAGGER_FILE}


.DEFAULT_GOAL := watch
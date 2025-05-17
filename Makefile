BINARY=convertapitos
CMD_DIR=./src/cmd
# SWAGGER
SWAGGER_DIR ?= ./src/cmd
SWAGGER_OPTIONS ?= --parseDependency --parseInternal --parseFuncBody --parseDepth 5


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
	swag init --dir ${SWAGGER_DIR} ${SWAGGER_OPTIONS}


swagger-gen:
ifndef FILE
	$(error FILE is required for swagger-gen. Example: make swagger-gen FILE=src/internal/user/handler.go)
endif
	swag init -g $(FILE) --parseInternal

.PHONY: swagger-fmt
swagger-fmt:
	swag fmt --dir ${SWAGGER_DIR}

.DEFAULT_GOAL := watch
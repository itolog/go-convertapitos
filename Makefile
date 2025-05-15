BINARY=convertapitos

# DEVELOPMENT
.PHONY: run
run:
	APP_ENV="development" go run ./src/cmd

.PHONY: run-gcf
run-gcf:
	go run -gcflags '-m -l' ./src/cmd

.PHONY: watch
watch:
	APP_ENV="development" air	
# BUILD
.PHONY: build
build:
	go build -o ./bin/${BINARY} ./src/cmd

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

.DEFAULT_GOAL := watch
BINARY=convertapitos

.PHONY: run
run:
	APP_ENV="development" go run ./src/cmd

.PHONY: run-gcf
run-gcf:
	go run -gcflags '-m -l' ./src/cmd

.PHONY: watch
watch:
	APP_ENV="development" air	

.PHONY: build
build:
	go build -o ./bin/${BINARY} ./src/cmd

.PHONY: run-prod
run-prod:
	APP_ENV="production" ./bin/${BINARY}

.PHONY: docker-up
docker-up:
	docker compose up -d --build

.PHONY: docker-stop
docker-stop:
	docker compose stop

.PHONY: docker-down
docker-down:
	docker compose down

.DEFAULT_GOAL := watch
BINARY=convertapitos
.PHONY: run watch build run-prod

run:
	APP_ENV="development" go run ./src/cmd

watch:
	APP_ENV="development" air	

build:
	go build -o ./bin/${BINARY} ./src/cmd

run-prod:
	APP_ENV="production" ./bin/${BINARY}

.DEFAULT_GOAL := watch
BINARY=convertapitos

.PHONY: run
run:
	APP_ENV="development" go run ./src/cmd

.PHONY: watch
watch:
	APP_ENV="development" air	

.PHONY: build
build:
	go build -o ./bin/${BINARY} ./src/cmd

.PHONY: run-prod
run-prod:
	APP_ENV="production" ./bin/${BINARY}

.DEFAULT_GOAL: run
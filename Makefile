SHELL=cmd.exe
FRONT_END_BINARY=frontend.exe
BROKER_BINARY=broker-service
LOGGER_BINARY=logger-service
AUTH_BINARY=auth-service

## up: starts all containers in the background without forcing build
up: start
	@echo Starting Docker images...
	docker-compose up -d
	@echo Docker images started!

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_auth build_logger
	@echo Stopping docker images (if running...)
	docker-compose down --volumes
	@echo Building (when required) and starting docker images...
	docker-compose up --build -d
	@echo Docker images built and started!

## down: stop docker compose
down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo Building broker binary...
	chdir ./broker-service && set GOOS=linux&&  set CGO_ENABLED=0&& go build -o ${BROKER_BINARY} ./cmd/api
	@echo Done!


## build_logger: builds the logger binary as a linux executable
build_logger:
	@echo Building logger binary...
	chdir ./logger-service && set GOOS=linux&&  set CGO_ENABLED=0&& go build -o ${LOGGER_BINARY} ./cmd/api
	@echo Done!

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo Building auth binary...
	chdir ./auth-service && set GOOS=linux&&  set CGO_ENABLED=0&& go build -o ${AUTH_BINARY} ./cmd/api
	@echo Done!

## build_front: builds the frontend binary
build_front:
	@echo Building front end binary...
	chdir ./frontend && go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo Done!

## start: starts the front end
start: build_front
	@echo Starting front end
	chdir ./frontend && start /B ${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo Stopping front end...
	@taskkill /IM "${FRONT_END_BINARY}" /F
	@echo "Stopped front end!"

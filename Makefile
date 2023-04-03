#@IgnoreInspection BashAddShebang

export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=form_server
export LDFLAGS="-w -s"

all: build test

build:
	go build -race -o ./build/$(APP) .

build-static:
	# CGO_ENABLED=0 go build -race -v -o ./build/$(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .
	go build -race -v -o ./build/$(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

run:
	PORT=3002 \
	ADDRESS="0.0.0.0" \
	DB_HOST=localhost \
	DB_PORT=5431 \
	DB_USERNAME=postgres \
	DB_PASSWORD='1234' \
	DB_NAME=zarinworld \
	go run -race .

############################################################
# Docker container
############################################################

container:
	docker build -t form:v1.0.1 .

run-container:
	docker run --rm -p 3000:3000 form:v1.0.1

############################################################
# Test
############################################################

test:
	go test -v -race ./...


.PHONY: build run build-static test container

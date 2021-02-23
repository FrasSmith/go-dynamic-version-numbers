dependencies:
	go mod tidy

build: dependencies
	go run cmd/buildnumber.go
	go build -o dynamic

run:
	./dynamic

all: build run
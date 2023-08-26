.PHONY: server web build build-osx

server:
	go run ./main.go
web:
	cd ./web && yarn dev
build:
	GOOS=linux CGO_ENABLED=0 go build -o ./ports-app ./main.go
build-osx:
	GOOS=darwin CGO_ENABLED=0 go build -0 ./build/ports-app ./main.go

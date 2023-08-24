.PHONY: server web build

server:
	go run ./main.go
web:
	cd ./web && yarn dev

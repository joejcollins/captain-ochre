# Ja ja

all: go-build

go-build:
	mkdir -p ./bin
	go build -o ./bin ./cmd/api/
	go build -o ./bin ./cmd/mkreports/ 

go-install:
	go install ./...


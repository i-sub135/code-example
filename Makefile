dbmigrate:
	go run migrate/*.go

start:
	go run src/*.go

build:
	go build src/*.go
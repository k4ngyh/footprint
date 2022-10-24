BINARY_NAME = footprint

build:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_NAME)-darwin-amd64  footprint.go
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux-amd64  footprint.go

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux

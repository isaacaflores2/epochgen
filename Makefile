BINARY_NAME=epochgen

build:
	go build -o bin/native/$(BINARY_NAME) main.go

build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/linux_x64/$(BINARY_NAME) main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/osx_x64/$(BINARY_NAME) main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/osx_arm64/$(BINARY_NAME) main.go
	GOOS=windows GOARCH=amd64 go build -o bin/win_x64/$(BINARY_NAME).exe main.go

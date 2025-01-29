

main_package_path = ./cmd/api
binary_name = main

tidy:
	go mod tidy -v
	go fmt ./...

build:
	go build -o=/tmp/bin/${binary_name} ${main_package_path}

run: build
	/tmp/bin/${binary_name}



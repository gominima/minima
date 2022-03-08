# minima is a free and open source software under MIT license
all: say_version build test clean

tests: build test clean

say_version:
	@echo "gominima/minima v1"

build:
	@echo "Building..."
	go build minima.go

test:
	@echo "Testing..."
	go test

clean:
	@echo "Cleaning up..."
	go fmt ./

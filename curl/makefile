build:
	go build -o bin/gocurl cmd/main.go

test:
	./bin/gocurl http://httpbin.org/get

clean:
	rm -rf bin

all: build test
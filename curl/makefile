build:
	go build -o bin/gocurl cmd/main.go

test:
	./bin/gocurl http://eu.httpbin.org:80/get

clean:
	rm -rf bin

all: build test
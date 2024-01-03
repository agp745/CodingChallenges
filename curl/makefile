build:
	go build -o bin/gocurl cmd/main.go

test:
	./bin/gocurl -X GET http://httpbin.org/get \
	-d '{"key": "value"}' \
	-H "Content-Type: application/json"

clean:
	rm -rf bin

all: build test
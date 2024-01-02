build:
	go build -o bin/gocurl cmd/main.go

test:
	./bin/gocurl -X DELETE http://httpbin.org/delete

clean:
	rm -rf bin

all: build test
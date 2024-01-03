build:
	go build -o bin/gocurl cmd/main.go

testGet:
	./bin/gocurl http://httpbin.org/get

testPost:
	./bin/gocurl -X POST http://httpbin.org/post \
	-d '{"key": "value"}' \
	-H "Content-Type: application/json"

testDelete:
	./bin/gocurl -X delete http://httpbin.org/delete

clean:
	rm -rf bin

get: build testGet

post: build testPost

delete: build testDelete
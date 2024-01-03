build:
	go build -o bin/gocurl cmd/main.go

testGet:
	./bin/gocurl http://httpbin.org/get

testDelete:
	./bin/gocurl -X delete -v http://httpbin.org/delete

testPost:
	./bin/gocurl -X POST http://httpbin.org/post \
	-d '{"key": "value"}' \
	-H "Content-Type: application/json"

testPut:
	./bin/gocurl -X PUT http://httpbin.org/put \
	-d '{"key2": "value2"}' \
	-H "Content-Type: application/json"

testPatch:
	./bin/gocurl -X PATCH http://httpbin.org/patch \
	-d '{"keyUpdate": "NEW VALUE"}' \
	-H "Content-Type: application/json"

testHead:
	./bin/gocurl -I http://httpbin.org/get

clean:
	rm -rf bin

get: build testGet

delete: build testDelete

post: build testPost

put: build testPut

patch: build testPatch

head: build testHead

build:
	go build -o bin/gocurl cmd/main.go

get:
	./bin/gocurl -K -v http://httpbin.org/get

delete:
	./bin/gocurl -K -X delete -v http://httpbin.org/delete

post:
	./bin/gocurl -X POST http://httpbin.org/post \
	-d '{"key": "value"}' \
	-H "Content-Type: application/json"

put:
	./bin/gocurl -X PUT http://httpbin.org/put \
	-d '{"key2": "value2"}' \
	-H "Content-Type: application/json"

patch:
	./bin/gocurl -X PATCH http://httpbin.org/patch \
	-d '{"keyUpdate": "NEW VALUE"}' \
	-H "Content-Type: application/json"

head:
	./bin/gocurl -I http://httpbin.org/get

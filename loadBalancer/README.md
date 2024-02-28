# Go Load Balancer

There are 2 main files:
- [lb](./cmd/lb/main.go) is the Load Balancer which implements a round robin algorithm, sending requests to 3 simple servers running on ports 70, 71, and 72
- [be](./cmd/be/main.go) is the server which simply responds to GET requests

## How to build
Execute ```make build``` to create the executables. It will be saved in the ```bin/``` directory.

## How to run
Execute three different web servers.

```bash
./bin/be -p 70
./bin/be -p 71
./bin/be -p 72
```
Now, execute the load balancer.
```bash
./bin/lb
```
You can use cURL to make a request to our load balancer.
```bash
curl http://localhost:8080
```

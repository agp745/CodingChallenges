package lb

import (
	"fmt"
	"net/http"
)

type LoadBalancer struct {
	port string
}

func NewLoadBalancer(port string) *LoadBalancer {
	return &LoadBalancer{
		port: port,
	}
}

func (lb *LoadBalancer) Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Recieved connection from %v\n%v / %v\nHost: %v\nUser-Agent: %v\nAccept: %v", r.RemoteAddr, r.Method, r.Proto, r.Host, r.UserAgent(), r.Header.Get("Accept"))
	})

	http.ListenAndServe(lb.port, nil)
}

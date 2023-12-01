package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

type LoadBalancer struct {
	port    int
	servers map[string]int
}

func NewLoadBalancer(port int) *LoadBalancer {
	servers := make(map[string]int)
	return &LoadBalancer{
		port:    port,
		servers: servers,
	}
}

func readClientRequest(r *http.Request) {
	fmt.Printf("Recieved connection from %v\n"+
		"%v / %v\n"+
		"Host: %v\n"+
		"User-Agent: %v\n"+
		"Accept: %v\n\n",
		r.RemoteAddr, r.Method, r.Proto, r.Host, r.UserAgent(), r.Header.Get("Accept"))
}

func modifyResponse() func(*http.Response) error {
	return func(res *http.Response) error {

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		fmt.Printf("Response from server: %v %v\n\n", res.Proto, res.Status)
		fmt.Println(string(body))

		res.Body = io.NopCloser(bytes.NewReader(body))
		res.ContentLength = int64(len(body))

		// Set the modified response headers
		res.Header.Set("Content-Length", strconv.Itoa(len(body)))

		return nil
	}
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	readClientRequest(r)

	targetUrl, err := url.Parse("http://localhost:70/server")
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	proxy.ModifyResponse = modifyResponse()
	proxy.ServeHTTP(w, r)
}

func (lb *LoadBalancer) Listen() {
	fmt.Printf("load balancer listening on http://localhost:%v\n\n", lb.port)

	http.HandleFunc("/", handleProxy)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(lb.port), nil))
}

func main() {
	lb := NewLoadBalancer(80)
	lb.servers["server1"] = 70
	lb.Listen()
}

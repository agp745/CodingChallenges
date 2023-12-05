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
	port           int
	servers        []string
	healthyServers []string
	badServers     []string
	currServerIdx  int
}

func NewLoadBalancer(port int) *LoadBalancer {
	servers := []string{":70", ":71", ":72"}
	return &LoadBalancer{
		port:          port,
		servers:       servers,
		currServerIdx: -1,
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
		// res.ContentLength = int64(len(body))
		// res.Header.Set("Content-Length", strconv.Itoa(len(body)))

		return nil
	}
}

func (lb *LoadBalancer) balanceServers() string {
	lb.currServerIdx += 1

	if lb.currServerIdx == len(lb.healthyServers) {
		lb.currServerIdx = 0
	}

	return "http://localhost" + lb.healthyServers[lb.currServerIdx]
}

func (lb *LoadBalancer) handleProxy(w http.ResponseWriter, r *http.Request) {
	readClientRequest(r)

	targetUrl, err := url.Parse(lb.balanceServers())
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

	http.HandleFunc("/", lb.handleProxy)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(lb.port), nil))
}

func main() {
	lb := NewLoadBalancer(80)
	go lb.HealthCheck()
	lb.Listen()
}

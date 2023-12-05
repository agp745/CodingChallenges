package main

import (
	"fmt"
	"net/http"
	"slices"
	"time"
)

func (lb *LoadBalancer) BadServer(port string) {

	if !slices.Contains(lb.badServers, port) {
		lb.badServers = append(lb.badServers, port)
	}

	if slices.Contains(lb.healthyServers, port) {
		idx := slices.Index(lb.healthyServers, port)
		lb.healthyServers = append(lb.healthyServers[:idx], lb.healthyServers[idx+1:]...)
	}
	// fmt.Printf("Unhealthy server on port %s\nResponded with: %v", port)
}

func (lb *LoadBalancer) HealthyServer(port string) {
	green := "\033[32m"
	reset := "\033[0m"

	if !slices.Contains(lb.healthyServers, port) {
		lb.healthyServers = append(lb.healthyServers, port)
	}

	if slices.Contains(lb.badServers, port) {
		idx := slices.Index(lb.badServers, port)
		lb.badServers = append(lb.badServers[:idx], lb.badServers[idx+1:]...)
	}

	fmt.Printf("%sHealthy server on port %s%s\n", green, port, reset)
}

func (lb *LoadBalancer) HealthCheck() {

	red := "\033[31m"
	bold := "\033[1m"
	reset := "\033[0m"

	for _, port := range lb.servers {
		resp, err := http.Get("http://localhost" + port)
		if err != nil {
			lb.BadServer(port)
			fmt.Printf("Server on port %s %s%sdown ❌%s\n", port, red, bold, reset)
			continue
		}

		if resp.StatusCode != 200 {
			lb.BadServer(port)
		}

		lb.HealthyServer(port)
	}
	fmt.Printf("healthy servers: %v\nbad servers: %v\n", lb.healthyServers, lb.badServers)
}

func (lb *LoadBalancer) TimedCheck() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("----------------")
		lb.HealthCheck()
	}

	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)
	// // <-interrupt

	// fmt.Println("Shutting down load balancer...")
}

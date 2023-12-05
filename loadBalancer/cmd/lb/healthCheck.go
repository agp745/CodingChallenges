package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
)

func (lb *LoadBalancer) BadServer(port string, status string) {

	if !slices.Contains(lb.badServers, port) {
		lb.badServers = append(lb.badServers, port)
	}

	if slices.Contains(lb.healthyServers, port) {
		idx := slices.Index(lb.healthyServers, port)
		lb.healthyServers = append(lb.healthyServers[:idx], lb.healthyServers[idx+1:]...)
	}
	fmt.Printf("Unhealthy server on port %s\nResponded with: %v", port, status)
}

func (lb *LoadBalancer) HealthyServer(port string) {
	if !slices.Contains(lb.healthyServers, port) {
		lb.healthyServers = append(lb.healthyServers, port)
	}
	fmt.Printf("Healthy server on port %s\n", port)
}

func (lb *LoadBalancer) HealthCheck() {
	for _, port := range lb.servers {
		resp, err := http.Get("http://localhost" + port)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != 200 {
			lb.BadServer(port, resp.Status)
		}

		lb.HealthyServer(port)
	}
}

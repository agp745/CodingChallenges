package main

import (
	"github.com/agp745/CodingChallenges/loadBalancer/lb"
)

func main() {
	lb := lb.NewLoadBalancer(":8080")

	lb.Start()
}

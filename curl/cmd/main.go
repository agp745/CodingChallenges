package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func SendReq(url *url.URL) {
	fmt.Printf("connecting to %s\n", url.Hostname())
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gocurl <URL>")
		return
	}

	url, err := url.Parse(strings.TrimSpace(os.Args[1]))
	if err != nil {
		panic(err)
	}

	SendReq(url)
}

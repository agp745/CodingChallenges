package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./bin/gocurl <URL>")
		return
	}

	u := ParseURL(os.Args[1])

	fmt.Fprintf(os.Stdout, "connecting to %s\r\n", u.Hostname())
	fmt.Fprintf(os.Stdout, "Sending request GET %s HTTP/1.1\r\n", u.Path)
	fmt.Fprintf(os.Stdout, "Host: %s\r\n", u.Hostname())
	fmt.Fprintln(os.Stdout, "Accept: */*\r")

}

func ParseURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	if len(u.Port()) == 0 {
		u.Host = u.Host + ":80"
	}

	return u
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func parseURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	if len(u.Port()) == 0 {
		u.Host = u.Host + ":80"
	}

	return u
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./bin/gocurl <URL>")
		return
	}

	u := parseURL(os.Args[1])

	client := &http.Client{}
	req := createRequest(u)
	sendRequest(client, req)
}

func createRequest(u *url.URL) *http.Request {
	req, err := http.NewRequest("GET", fmt.Sprint(u), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Host", u.Hostname())
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "close")

	fmt.Fprintf(os.Stdout, "connecting to %s\r\n", u.Hostname())
	fmt.Fprintf(os.Stdout, "Sending request %s %s %s\r\n", req.Method, u.Path, req.Proto)
	fmt.Fprintf(os.Stdout, "Host: %s\r\n", u.Hostname())
	fmt.Fprintf(os.Stdout, "Accept: %s\r\n\n", req.Header.Get("Accept"))

	return req
}

func sendRequest(client *http.Client, req *http.Request) {
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	conn := res.Header.Get("Connection")
	if len(conn) == 0 {
		conn = "close"
	}

	fmt.Fprintf(os.Stdout, "%s %s\n", res.Proto, res.Status)
	fmt.Fprintf(os.Stdout, "Date: %s\n", res.Header.Get("Date"))
	fmt.Fprintf(os.Stdout, "Content-Type: %s\n", res.Header.Get("Content-Type"))
	fmt.Fprintf(os.Stdout, "Content-Length: %s\n", res.Header.Get("Content-Length"))
	fmt.Fprintf(os.Stdout, "Connection: %s\n", conn)
	fmt.Fprintf(os.Stdout, "Server: %s\n", res.Header.Get("Server"))
	fmt.Fprintf(os.Stdout, "Access-Control-Allow-Origin: %s\n", res.Header.Get("Access-Control-Allow-Origin"))
	fmt.Fprintf(os.Stdout, "Access-Control-Allow-Credentials: %s\n\n", res.Header.Get("Access-Control-Allow-Credentials"))

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

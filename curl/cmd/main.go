package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type GocurlClient struct {
	httpClient  *http.Client
	httpRequest *http.Request
	verbose     bool
}

func initGocurlClient() *GocurlClient {
	client := &http.Client{}
	return &GocurlClient{
		httpClient: client,
		// verbose:    true,
	}
}

func initFlags() *bool {
	return flag.Bool("v", false, "prints the request and response headers to stdout")
}

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
	gc := initGocurlClient()
	vFlag := initFlags()
	var rawUrl string

	flag.Parse()

	switch len(os.Args) {
	case 2:
		rawUrl = os.Args[1]
	case 3:
		rawUrl = os.Args[2]
	default:
		fmt.Println("Usage: ./bin/gocurl <URL>")
		return
	}

	if *vFlag {
		gc.verbose = true
	}

	u := parseURL(rawUrl)

	gc.createRequest(u)
	gc.sendRequest()
}

func (gc *GocurlClient) createRequest(u *url.URL) {
	prefix := ">"
	req, err := http.NewRequest("GET", fmt.Sprint(u), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Host", u.Hostname())
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "close")

	gc.httpRequest = req

	if gc.verbose {
		fmt.Fprintf(os.Stdout, "%s connecting to %s\r\n", prefix, u.Hostname())
		fmt.Fprintf(os.Stdout, "%s Sending request %s %s %s\r\n", prefix, req.Method, u.Path, req.Proto)
		fmt.Fprintf(os.Stdout, "%s Host: %s\r\n", prefix, u.Hostname())
		fmt.Fprintf(os.Stdout, "%s Accept: %s\r\n", prefix, req.Header.Get("Accept"))
		fmt.Fprintf(os.Stdout, "%s\r\n", prefix)
	}
}

func (gc *GocurlClient) sendRequest() {
	prefix := "<"
	res, err := gc.httpClient.Do(gc.httpRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	conn := res.Header.Get("Connection")
	if len(conn) == 0 {
		conn = "close"
	}

	if gc.verbose {
		fmt.Fprintf(os.Stdout, "%s %s %s\n", prefix, res.Proto, res.Status)
		fmt.Fprintf(os.Stdout, "%s Date: %s\n", prefix, res.Header.Get("Date"))
		fmt.Fprintf(os.Stdout, "%s Content-Type: %s\n", prefix, res.Header.Get("Content-Type"))
		fmt.Fprintf(os.Stdout, "%s Content-Length: %s\n", prefix, res.Header.Get("Content-Length"))
		fmt.Fprintf(os.Stdout, "%s Connection: %s\n", prefix, conn)
		fmt.Fprintf(os.Stdout, "%s Server: %s\n", prefix, res.Header.Get("Server"))
		fmt.Fprintf(os.Stdout, "%s Access-Control-Allow-Origin: %s\n", prefix, res.Header.Get("Access-Control-Allow-Origin"))
		fmt.Fprintf(os.Stdout, "%s Access-Control-Allow-Credentials: %s\n", prefix, res.Header.Get("Access-Control-Allow-Credentials"))
		fmt.Fprintf(os.Stdout, "%s\r\n", prefix)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

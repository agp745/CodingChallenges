package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"slices"
	"strings"
)

type GocurlClient struct {
	httpClient    *http.Client
	httpRequest   *http.Request
	requestMethod string
	verbose       bool
}

const (
	RED   = "\033[31m"
	RESET = "\033[0m"
	BOLD  = "\033[1m"
)

func initGocurlClient() *GocurlClient {
	client := &http.Client{}
	return &GocurlClient{
		httpClient: client,
	}
}

func initFlags() (*bool, *string) {
	v := flag.Bool("v", false, "Return request and response headers to stdout")
	X := flag.String("X", "GET", "Set the request method. Default: GET")

	return v, X
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

func checkMethod(m string) bool {
	methods := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	return slices.Contains(methods, m)
}

func main() {
	gc := initGocurlClient()
	vFlag, XFlag := initFlags()
	var rawUrl string

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("Usage: ./bin/gocurl [flags] <URL>")
		return
	}
	rawUrl = os.Args[len(os.Args)-1]

	if *vFlag {
		gc.verbose = true
	}
	method := strings.ToUpper(strings.TrimSpace(*XFlag))
	if ok := checkMethod(method); !ok {
		fmt.Printf("%sError:%s %s%s%s is NOT a valid method\r\n", RED, RESET, BOLD, method, RESET)
		return
	}
	gc.requestMethod = method

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
	req.Method = gc.requestMethod

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

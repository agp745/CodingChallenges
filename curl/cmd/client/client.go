package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	HttpClient    *http.Client
	HttpRequest   *http.Request
	URL           *url.URL
	RequestMethod string
	RequestBody   *string
	// requestHeaders *string
	Verbose bool
}

func NewClient() *Client {
	client := &http.Client{}
	return &Client{
		HttpClient: client,
	}
}

// func (gc *Client) CreateRequest(u *url.URL) {
func (gc *Client) CreateRequest() {
	prefix := ">"
	u := gc.URL

	req, err := http.NewRequest("GET", fmt.Sprint(u), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Host", u.Hostname())
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "close")
	req.Method = gc.RequestMethod

	if len(*gc.RequestBody) > 0 {
		req.Body.Read([]byte(*gc.RequestBody))
	}

	gc.HttpRequest = req

	if gc.Verbose {
		fmt.Fprintf(os.Stdout, "%s connecting to %s\r\n", prefix, u.Hostname())
		fmt.Fprintf(os.Stdout, "%s Sending request %s %s %s\r\n", prefix, req.Method, u.Path, req.Proto)
		fmt.Fprintf(os.Stdout, "%s Host: %s\r\n", prefix, u.Hostname())
		fmt.Fprintf(os.Stdout, "%s Accept: %s\r\n", prefix, req.Header.Get("Accept"))
		fmt.Fprintf(os.Stdout, "%s\r\n", prefix)
	}
}

func (gc *Client) SendRequest() {
	prefix := "<"
	res, err := gc.HttpClient.Do(gc.HttpRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	conn := res.Header.Get("Connection")
	if len(conn) == 0 {
		conn = "close"
	}

	if gc.Verbose {
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

package client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	BOLD  = "\033[1m"
	RESET = "\033[0m"
)

type Client struct {
	HttpClient     *http.Client
	HttpRequest    *http.Request
	URL            *url.URL
	RequestMethod  string
	RequestBody    *string
	RequestHeaders map[string]string
	Verbose        bool
	HeadRequest    bool
	KeepAlive      bool
}

func NewClient(v bool, I bool, K bool) *Client {
	client := &http.Client{}
	return &Client{
		HttpClient:  client,
		Verbose:     v,
		HeadRequest: I,
		KeepAlive:   K,
	}
}

func (gc *Client) headRequest() {
	res, err := gc.HttpClient.Head(fmt.Sprint(gc.URL))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Proto, res.Status)
	for k, v := range res.Header {
		fmt.Printf("%s%s:%s %s\r\n", BOLD, k, RESET, v)
	}
}

func (gc *Client) CreateRequest() {
	prefix := ">"
	u := gc.URL

	if gc.HeadRequest {
		gc.headRequest()
		return
	}

	req, err := http.NewRequest(gc.RequestMethod, fmt.Sprint(u), bytes.NewBuffer([]byte(*gc.RequestBody)))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Host", u.Hostname())
	req.Header.Set("Accept", "*/*")
	if gc.KeepAlive {
		req.Header.Set("Connection", "keep-alive")
	} else {
		req.Header.Set("Connection", "close")
	}
	for k, v := range gc.RequestHeaders {
		req.Header.Set(k, v)
	}

	gc.HttpRequest = req

	if gc.Verbose {
		fmt.Println(prefix, req.Method, req.URL.Path, req.Proto)
		for k, v := range req.Header {
			fmt.Printf("%s %s: %s \r\n", prefix, k, v[0])
		}
		fmt.Println(prefix)
	}

	gc.SendRequest()
}

func (gc *Client) SendRequest() {
	prefix := "<"

	res, err := gc.HttpClient.Do(gc.HttpRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if gc.Verbose {
		fmt.Println(prefix, res.Proto, res.Status)
		for k, v := range res.Header {
			fmt.Printf("%s %s: %s\r\n", prefix, k, v[0])
		}
		fmt.Println(prefix)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

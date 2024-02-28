package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"slices"
	"strings"

	"github.com/agp745/CodingChallenges/curl/cmd/client"
)

const (
	RED   = "\033[31m"
	RESET = "\033[0m"
	BOLD  = "\033[1m"
)

func initFlags() (*bool, *string, *string, *string, *bool, *bool) {
	// VERBOSE
	v := flag.Bool("v", false, "Return request and response headers to stdout")
	// METHOD
	X := flag.String("X", "GET", "Set the request method. Usage: ./bin/gocurl -X GET <URL>")
	// BODY
	d := flag.String("d", "", "Set request body")
	// HEADERS
	H := flag.String("H", "", "Set request headers")
	// HEAD
	I := flag.Bool("I", false, "Get Response Headers without requesting the body")
	// KEEP-ALIVE
	K := flag.Bool("K", false, "sets Connection header to Keep-Alive")
	return v, X, d, H, I, K
}

func checkMethod(m string) bool {
	methods := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	return slices.Contains(methods, m)
}

func getUrl() *url.URL {
	if len(os.Args) == 1 {
		fmt.Println("Usage: ./bin/gocurl [flags] <URL>")
		os.Exit(0)
	}

	for _, arg := range os.Args {
		rawUrl := arg
		if !strings.Contains(rawUrl, "://") {
			rawUrl = fmt.Sprintf("https://%s", arg)
		}
		u, err := url.ParseRequestURI(rawUrl)
		fmt.Println("YER")
		if err != nil {
			continue
		}

		if len(u.Port()) == 0 {
			switch u.Scheme {
			case "http":
				u.Host = u.Host + ":80"
			case "https":
				u.Host = u.Host + ":443"
			}
		}
		return u
	}

	return nil
}

func main() {
	vFlag, XFlag, dFlag, HFlag, IFlag, KFlag := initFlags()
	flag.Parse()
	// Parse flags after first argument
	if len(flag.Args()) > 1 {
		for i, arg := range flag.Args() {
			if arg == "-d" || arg == "--d" {
				dFlag = &flag.Args()[i+1]
			}
			if arg == "-H" || arg == "--H" {
				HFlag = &flag.Args()[i+1]
			}
		}
	}

	gc := client.NewClient(*vFlag, *IFlag, *KFlag)
	gc.URL = getUrl()

	method := strings.ToUpper(strings.TrimSpace(*XFlag))
	if ok := checkMethod(method); !ok {
		fmt.Printf("%sError:%s %s%s%s is NOT a valid method\r\n", RED, RESET, BOLD, method, RESET)
		return
	}
	gc.RequestMethod = method
	gc.RequestBody = dFlag
	if len(*HFlag) > 0 {
		gc.RequestHeaders = parseReqHeaders(*HFlag)
	}

	gc.CreateRequest()
}

func parseReqHeaders(h string) map[string]string {
	headers := make(map[string]string)

	arr := strings.Split(h, ": ")

	for i := 0; i < len(arr); i += 2 {
		headers[arr[i]] = arr[i+1]
	}

	return headers
}

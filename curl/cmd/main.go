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

func initFlags() (*bool, *string, *string, *string) {
	// VERBOSE
	v := flag.Bool("v", false, "Return request and response headers to stdout")
	// METHOD
	X := flag.String("X", "GET", "Set the request method. Usage: ./bin/gocurl -X GET <URL>")
	// BODY
	d := flag.String("d", "", "Set request body")
	// // HEADERS
	H := flag.String("H", "", "Set request headers")
	return v, X, d, H
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
		u, err := url.ParseRequestURI(arg)
		if err != nil {
			continue
		}

		if len(u.Port()) == 0 {
			u.Host = u.Host + ":80"
		}
		return u
	}

	return nil
}

func main() {
	vFlag, XFlag, dFlag, HFlag := initFlags()
	gc := client.NewClient()
	gc.URL = getUrl()

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

	if *vFlag {
		gc.Verbose = true
	}

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

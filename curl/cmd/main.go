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

func initFlags() (*bool, *string, *string) {
	// VERBOSE
	v := flag.Bool("v", false, "Return request and response headers to stdout")
	// METHOD
	X := flag.String("X", "GET", "Set the request method. Usage: ./bin/gocurl -X GET <URL>")
	// BODY
	d := flag.String("d", "", "Set request body")
	// // HEADERS
	// H := flag.String("H", "", "Set request headers")
	return v, X, d
}

func checkMethod(m string) bool {
	methods := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	return slices.Contains(methods, m)
}

// func parseURL(s string) *url.URL {
// 	u, err := url.Parse(s)
// 	if err != nil {
// 		return nil
// 	}

// 	if len(u.Port()) == 0 {
// 		u.Host = u.Host + ":80"
// 	}

// 	return u
// }

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
	vFlag, XFlag, dFlag := initFlags()
	flag.Parse()
	gc := client.NewClient()
	gc.URL = getUrl()

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
	fmt.Println("D FLAG", *dFlag)

	gc.CreateRequest()
	gc.SendRequest()
}

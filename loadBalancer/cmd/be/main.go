package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Backend struct {
	port int
}

func main() {
	server := NewServer(70)
	server.Listen()
}

func NewServer(port int) *Backend {
	return &Backend{
		port: port,
	}
}

func connHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Recieved connection from %v\n"+
		"%v / %v\n"+
		"Host: %v\n"+
		"User-Agent: %v\n"+
		"Accept: %v\n\n"+
		"Replied with a hello message\n\n",
		r.RemoteAddr, r.Method, r.Proto, r.Host, r.UserAgent(), r.Header.Get("Accept"))

	w.Write([]byte("Hello from Backend Server"))
}

func (b *Backend) Listen() {
	fmt.Printf("server listening on http://localhost:%v\n\n", b.port)

	http.HandleFunc("/", connHandler)

	if err := http.ListenAndServe(":"+strconv.Itoa(b.port), nil); err != nil {
		log.Fatal(err)
	}
}

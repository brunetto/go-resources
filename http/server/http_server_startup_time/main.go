package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

func myHandler(w http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(w, "hello, world!\n")
}

func main() {
	t0 := time.Now()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(myHandler),
	}

	ln, err := net.Listen("tcp", ":8080")
	dieIf(err)

	lst := Listener{t0: t0, once: &sync.Once{}}
	lst.Listener = ln

	log.Fatal(srv.Serve(lst))
}

type Listener struct {
	t0 time.Time
	net.Listener
	once *sync.Once
}

func (l Listener) Accept() (net.Conn, error) {
	l.once.Do(func() {
		fmt.Println("Server ready with startup time",
			time.Since(l.t0))
	})

	return l.Listener.Accept()
}

func dieIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

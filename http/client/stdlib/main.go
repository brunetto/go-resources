package main

import (
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

func main() {
	c := &http.Client{
		Timeout: 3 * time.Second,
		// Transport: <your custom http.RoundTripper here>,
	}

	// same as
	// req, err := http.NewRequest("GET", "https://example.com", nil)
	// if err != nil {
	// 	log.Fatal(errors.Wrap(err, "can0t create new request"))
	// }
	// rsp, err := c.Do(req)

	rsp, err := c.Get("https://example.com")
	if err != nil {
		log.Fatal(errors.Wrap(err, "can't get example.com"))
	}
	defer rsp.Body.Close()

	if rsp.StatusCode < http.StatusOK || rsp.StatusCode >= http.StatusBadRequest {
		log.Fatalf("response status is %v - %v", rsp.StatusCode, rsp.Status)
	}

	// ...
}

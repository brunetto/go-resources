package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = w.Write([]byte("Hello again."))
	})

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		// --------------------------------
		// other custom quantities?
		// MaxHeaderBytes: 1 << 20,
		//---------------------------------
		// https
		// Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// 	w.Header().Set("Connection", "close")
		// 	url := "https://" + req.Host + req.URL.String()
		// 	http.Redirect(w, req, url, http.StatusMovedPermanently)
		// }),
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
	// log.Fatal(srv.ListenAndServeTLS())
}

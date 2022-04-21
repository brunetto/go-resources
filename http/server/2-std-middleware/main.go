package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// here accepting and returning http.HandlerFunc automatically cast
// func(http.ResponseWriter, *http.Request) to http.Handler
// allowing us to pass the function as a Handler
func withTimedLogger(lg zerolog.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()
		lg.Info().Msg("got new request")
		defer func() { lg.Info().Msg(fmt.Sprintf("handled in %v", time.Since(t0))) }()

		next.ServeHTTP(w, r)
	}
}

func hfn(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hi there!!"))
}

func main() {
	lg := zerolog.New(os.Stderr).With().Timestamp().Logger()

	mux := http.NewServeMux()
	mux.Handle("/", withTimedLogger(lg, hfn)) // you can also have a global middlewere, 
						  // see https://eli.thegreenplace.net/2021/rest-servers-in-go-part-5-middleware/

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":8080",
		Handler:      mux,
	}

	log.Fatal().Msg(srv.ListenAndServe().Error())
}

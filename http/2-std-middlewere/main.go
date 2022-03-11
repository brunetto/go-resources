package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

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
	mux.HandleFunc("/", withTimedLogger(lg, hfn))

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":8080",
		Handler:      mux,
	}

	log.Fatal().Msg(srv.ListenAndServe().Error())
}

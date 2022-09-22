package hhttp

import (
	"bytes"
	"io"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/brunetto/go-resources/lambda/helpers/hctx"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type StatusRecorder struct {
	StatusCode int
	Response   *bytes.Buffer
	http.ResponseWriter
	statusCodeWritten bool
}

func (sr *StatusRecorder) WriteHeader(statuscode int) {
	sr.statusCodeWritten = true
	sr.StatusCode = statuscode
	sr.ResponseWriter.WriteHeader(statuscode)
}

// Write is needed to wrap the original write function otherwise our WriteHeader won't be called
// https://stackoverflow.com/questions/66367237/cant-track-http-response-code-in-middleware
func (sr *StatusRecorder) Write(p []byte) (int, error) {
	if !sr.statusCodeWritten {
		sr.WriteHeader(http.StatusOK)
	}

	if sr.StatusCode != http.StatusOK {
		// record error response
		return io.MultiWriter(sr.Response, sr.ResponseWriter).Write(p)
	}

	return sr.ResponseWriter.Write(p)
}

func NewLogAndRecover(lg zerolog.Logger) func(next http.Handler) http.Handler { // middleware constructor
	return func(next http.Handler) http.Handler { // middleware, for each call will return
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // contextual logging wrapper
			t0 := time.Now()

			// get context amd enrich logger (e.g. with request id)
			ctx := r.Context()
			lg := hctx.CtxLogger(ctx, lg)

			// wrap http response
			sr := &StatusRecorder{Response: &bytes.Buffer{}, ResponseWriter: w}

			defer func() { // log response status and duration
				lg2 := lg.With().
					Str("path", r.URL.Path).
					Str("method", r.Method).
					Int("status_code", sr.StatusCode).
					Float64("duration_ms", float64(time.Since(t0).Milliseconds())).
					Float64("duration_ns", float64(time.Since(t0).Nanoseconds())).
					Logger()

				if sr.StatusCode < http.StatusOK || sr.StatusCode > http.StatusBadRequest {
					lg2.Error().Err(errors.New(sr.Response.String())).Send()
				} else {
					lg2.Info().Send()
				}
			}()

			defer func() { // recover panic and log it properly, just in case
				if r := recover(); r != nil {
					lg.Error().Fields(map[string]interface{}{"trace": debug.Stack()}).
						Err(errors.Errorf("recovered paninc: %v", r)).Send()

					http.Error(sr, "", http.StatusInternalServerError)
				}
			}()

			// pass contextual logger in the request
			lgr := lg.With().Logger()
			lctx := (&lgr).WithContext(ctx)
			lr := r.WithContext(lctx)

			// call next middleware/handler in the chain
			next.ServeHTTP(sr, lr)
		})
	}
}

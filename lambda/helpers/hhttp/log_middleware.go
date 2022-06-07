package hhttp

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/brunetto/go-resources/lambda/helpers/hctx"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type StatusRecorder struct {
	StatusCode int
	http.ResponseWriter
}

func (sr *StatusRecorder) WriteHeader(statuscode int) {
	sr.StatusCode = statuscode
	sr.ResponseWriter.WriteHeader(statuscode)
}

func NewLogAndRecover(lg zerolog.Logger) func(next http.Handler) http.Handler { // middleware constructor
	return func(next http.Handler) http.Handler { // middleware, for each call will return
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // contextual logging wrapper
			t0 := time.Now()

			// get context amd enrich logger (e.g. with request id)
			ctx := r.Context()
			lg = hctx.CtxLogger(ctx, lg)

			// wrap http response
			sr := &StatusRecorder{ResponseWriter: w}

			defer func() { // log response status and duration
				lg2 := lg.With().
					Int("code", sr.StatusCode).
					Float64("duration_ms", float64(time.Since(t0).Milliseconds())).
					Logger()

				if sr.StatusCode == http.StatusOK {
					lg2.Info().Send()
				} else {
					lg2.Error().Send()
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

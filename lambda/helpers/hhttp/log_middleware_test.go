package hhttp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestStatusRecorder_WriteHeader(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
	}{
		{name: "ok", statusCode: http.StatusOK},
		{name: "not found", statusCode: http.StatusNotFound},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			sr := &StatusRecorder{ResponseWriter: w}

			sr.WriteHeader(tt.statusCode)

			assert.Equal(t, tt.statusCode, sr.StatusCode)
			assert.Equal(t, tt.statusCode, w.Code)
		})
	}
}

func TestNewLogAndRecover(t *testing.T) {
	tests := []struct {
		name                string
		logRecorder         *bytes.Buffer
		handler             http.HandlerFunc
		wantStatus          int
		wantLevel           string
		wantResponseMessage string
	}{
		{
			name:                "ok",
			logRecorder:         &bytes.Buffer{},
			handler:             func(w http.ResponseWriter, req *http.Request) { _, _ = w.Write([]byte("OK")) },
			wantStatus:          http.StatusOK,
			wantLevel:           "info",
			wantResponseMessage: "OK",
		},
		{
			name:        "not found",
			logRecorder: &bytes.Buffer{},
			handler:     func(w http.ResponseWriter, req *http.Request) { w.WriteHeader(http.StatusNotFound) },
			wantStatus:  http.StatusNotFound,
			wantLevel:   "error",
		},
		{
			name:        "internal server error",
			logRecorder: &bytes.Buffer{},
			handler: func(w http.ResponseWriter, req *http.Request) {
				http.Error(w, "test error", http.StatusInternalServerError)
			},
			wantStatus:          http.StatusInternalServerError,
			wantLevel:           "error",
			wantResponseMessage: "test error\n",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			lg := zerolog.New(tt.logRecorder)
			middleware := NewLogAndRecover(lg)

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()

			router := mux.NewRouter()
			router.Use(middleware)
			router.HandleFunc("/", tt.handler)
			router.ServeHTTP(w, r)

			rsp := w.Result()
			defer rsp.Body.Close()

			var msg struct {
				StatusCode int    `json:"status_code"`
				Level      string `json:"level"`
				Message    string `json:"message"`
			}
			err := json.NewDecoder(tt.logRecorder).Decode(&msg)
			if !assert.Nil(t, err) {
				return
			}

			assert.Equal(t, tt.wantStatus, msg.StatusCode)
			assert.Equal(t, tt.wantLevel, msg.Level)

			responseMessage, err := io.ReadAll(rsp.Body)
			if !assert.Nil(t, err) {
				return
			}

			assert.Equal(t, tt.wantResponseMessage, string(responseMessage))
		})
	}
}

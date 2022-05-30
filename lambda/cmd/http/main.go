package main

import (
	"encoding/json"
	"net/http"
	"os"

	"go-resources/lambda/example"
	"go-resources/lambda/helpers/hhttp"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
)

func main() {
	lg := zerolog.New(os.Stdout)
	defer func() { lg.Info().Msg("Shutting down lambda instance") }()

	lg.Info().Msg("Starting lambda instance")

	//
	// your setup here: aws session, connections etc
	//

	// set-up http router for gorilla/mux proxy:
	// * no real need, you can do without it, it's just a helper
	//		to have the same signature as a http handler
	// * there're equivalent for net/http, gin-gonic, ...
	router := mux.NewRouter()

	router.Use(hhttp.NewLogAndRecover(lg)) // logging middleware
	router.Path("/usage").
		Queries("query_param_1", "{query_param_1}", "query_param_2", "{query_param_2}").
		Methods("POST").
		Handler(newHandler( /* your depts here */ ))

	// ⬆️ Everything above here is called once per instance, not once per execution (invocation).
	// This means we don't waste time initializing things every time.

	// ⬇️ Everything below here is called for each execution (invocation).
	lambda.Start(gorillamux.New(router))
}

func newHandler( /* your depts here */ ) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		lg := zerolog.Ctx(ctx) // retrieve contextual logger from request

		// get and validate params from requests
		vars := mux.Vars(req)
		qp1, exists := vars["query_param_1"]
		if !exists {
			http.Error(w, errors.Errorf("query param 1 not provided").Error(), http.StatusBadRequest)
			return
		}

		// just logging and not using query param this time 🤷‍♀️
		lg.Info().Str("qp1", qp1).Msg("got input query params")

		// Decode request
		var in example.In
		err := json.NewDecoder(req.Body).Decode(&in)
		if err != nil {
			http.Error(w, errors.Wrap(err, "can't json decode input json").Error(), http.StatusBadRequest)
			return
		}

		// Execute business logic
		out, err := example.DoStuff(ctx, in)
		if err != nil {
			http.Error(w, errors.Wrap(err, "can't do stuff").Error(), http.StatusBadRequest)
			return
		}

		// Encode response
		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			http.Error(w, errors.Wrap(err, "can't json encode response").Error(), http.StatusInternalServerError)
			return
		}
	}
}

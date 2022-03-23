* [sdtblib](./stdlib/)
    * don't create a HTTP client interface, `http.Client` is already a wrapper around the `http.Roundtripper` interface, implement that (see https://github.com/googleapis/google-api-go-client/issues/146), so that you can pass a concrete `http.Client` around.
* [Resty v2](https://github.com/go-resty/resty)
* [Request](https://github.com/monaco-io/request)
* [Heimdall](https://github.com/gojek/heimdall)
* [Clean HTTP](https://github.com/hashicorp/go-cleanhttp)

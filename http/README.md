# Basics

* `http.Handler` is an interface (with a single method `ServeHTTP(ResponseWriter, *Request)`), whose implementation is used by the server to serve requests.
* `http.HandleFunc` is a type allowing a fingle function `func(ResponseWriter, *Request)` to be used as handler
* `http.ServerMux` is a multiplexer (and itself a `http.Handler`), a router which binds a pattern (route) to a handler

# Suggested path

1. [stdlib-0](./0-basic/), [stdlib-1](./1-std/), [stdlib-2](./2-std-middlewere/)
2. [gorilla/mux](https://github.com/gorilla/mux)
3. [Gin Gonic](https://github.com/gin-gonic/gin)
4. [Buffalo](https://gobuffalo.io/it/)
5. Extra: [Ardan Labs/Bill Kennedy service code, talks and workshops](https://github.com/ardanlabs/service) as inspiration and [go-kit](https://gokit.io/) as reading

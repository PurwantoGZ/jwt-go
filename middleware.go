package main

import "net/http"

//CustomMux mux for middleware
type CustomMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

//RegisterMiddleware for register middleware
func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

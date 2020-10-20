package middlewares

import (
	"net/http"
)

func ProductsMiddlewareHandler(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		handler.ServeHTTP(w, r)
	})
}
// Package apigateway provides a gate way to access data via an api, it self does not contain any direct database access or buusiness logic.
package apigateway

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type HTTPRouting map[string]map[string]RestHandler

func NewHTTPRouter(routing HTTPRouting) *chi.Mux {
	r := chi.NewRouter()

	for path, route := range routing {
		for method, handler := range route {
			r.Method(method, path, http.HandlerFunc(handler))
		}
	}

	return r
}

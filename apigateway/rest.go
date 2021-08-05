package apigateway

import "net/http"

type RestHandler func(http.ResponseWriter, *http.Request)

func REST() HTTPRouting {
	endpoint := HTTPRouting{}

	endpoint["/inventory"]["get"] = func(http.ResponseWriter, *http.Request) {}

	return endpoint
}

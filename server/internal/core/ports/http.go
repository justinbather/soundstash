package ports

import "net/http"

type HttpHandler interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)

	Download(w http.ResponseWriter, r *http.Request)
}

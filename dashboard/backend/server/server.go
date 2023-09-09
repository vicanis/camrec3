package server

import (
	"dashboard/server/api"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() error {
	mx := mux.NewRouter()

	api.Configure(mx.PathPrefix("/api").Subrouter())

	mx.NotFoundHandler = http.FileServer(http.Dir("public"))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mx,
	}

	return srv.ListenAndServe()
}

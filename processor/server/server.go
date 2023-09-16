package server

import (
	"net/http"
	"processor/server/api"

	"github.com/gorilla/mux"
)

func Start() error {
	mx := mux.NewRouter()

	api.Configure(mx.PathPrefix("/api").Subrouter())

	srv := &http.Server{
		Addr:    ":80",
		Handler: mx,
	}

	return srv.ListenAndServe()
}

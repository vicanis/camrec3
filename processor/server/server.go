package server

import (
	"net/http"
	"processor/server/api"
	"time"

	"github.com/gorilla/mux"
)

func Start() error {
	mx := mux.NewRouter()

	api.Configure(mx.PathPrefix("/api").Subrouter())

	srv := &http.Server{
		Addr:         ":80",
		Handler:      mx,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return srv.ListenAndServe()
}

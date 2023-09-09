package api

import (
	"dashboard/database"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Configure(mx *mux.Router) {
	mx.Path("/").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte("not implemented"))
		},
	)

	mx.Path("/events").Methods(http.MethodGet).Handler(JsonResponse(
		func(r *http.Request) (any, error) {
			return database.GetItems()
		},
	))
}

func JsonResponse(handler func(r *http.Request) (any, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := handler(r)

		if err != nil {
			errorResponse := JsonError{
				Error: err.Error(),
			}

			jdata, jerr := json.Marshal(errorResponse)
			if jerr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jdata)

			return
		}

		buf, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	})
}

type JsonError struct {
	Error string `json:"error"`
}

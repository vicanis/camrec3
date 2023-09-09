package api

import (
	"dashboard/database"
	"encoding/json"
	"errors"
	"net/http"
	"time"

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
			arg := r.URL.Query().Get("day")
			if arg == "" {
				return nil, errors.New("no day request argument")
			}

			day, err := time.Parse("20060102", arg)
			if err != nil {
				return nil, err
			}

			return database.GetItems(day)
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

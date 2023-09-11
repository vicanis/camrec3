package server

import (
	"fmt"
	"log"
	"net/http"
	"processor/bundler"
	"time"

	"github.com/gorilla/mux"
)

func Start() error {
	mx := mux.NewRouter()

	mx.Path("/play/{timestamp}").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		timestamp := vars["timestamp"]

		if err := renderPlayer(fmt.Sprintf("/event/%s", timestamp), w); err != nil {
			log.Printf("template render failed: %s", err)
		}
	})

	mx.Path("/event/{timestamp}").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		timestamp := vars["timestamp"]

		if timestamp == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no timestamp"))
			return
		}

		ts, err := time.ParseInLocation("20060102150405", timestamp, time.Local)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid timestamp: %s", timestamp)
			return
		}

		data, err := bundler.SearchVideoBundle(ts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error searching files")
		}

		w.Header().Add("Content-Type", "video/mp4")
		w.Header().Add("Content-Length", fmt.Sprint(len(data)))

		w.WriteHeader(http.StatusOK)

		w.Write(data)
	})

	srv := &http.Server{
		Addr:    ":8088",
		Handler: mx,
	}

	return srv.ListenAndServe()
}

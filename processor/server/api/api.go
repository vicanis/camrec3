package api

import (
	"fmt"
	"log"
	"net/http"
	"processor/bundler"
	"time"

	"github.com/gorilla/mux"
)

func Configure(mx *mux.Router) {
	mx.Path("/play/{timestamp}").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		timestamp := vars["timestamp"]

		if err := renderPlayer(fmt.Sprintf("/api/event/%s", timestamp), w); err != nil {
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
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Add("Content-Type", "video/mp4")
		w.Header().Add("Content-Length", fmt.Sprint(len(data)))

		w.WriteHeader(http.StatusOK)

		w.Write(data)
	})
}

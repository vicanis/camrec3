package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"processor/bundler"
	"processor/encoder"
	"time"

	"github.com/gorilla/mux"
)

func Configure(mx *mux.Router) {
	mx.Path("/play/{timestamp}").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts, err := parseTimestamp(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
			return
		}

		if err := renderPlayer(fmt.Sprintf("/api/event/%s", ts.Format("20060102150405")), w); err != nil {
			log.Printf("template render failed: %s", err)
		}
	})

	mx.Path("/event/{timestamp}").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts, err := parseTimestamp(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
			return
		}

		data, err := bundler.SearchVideoBundle(ts)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		chunkStart := ts.Add(-3 * time.Second)
		_, _, s := chunkStart.Clock()

		data, err = encoder.Encode(data, s)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "encode failed: %s", err)
			return
		}

		w.Header().Add("Content-Type", "video/mp4")
		w.Header().Add("Content-Length", fmt.Sprint(len(data)))

		w.WriteHeader(http.StatusOK)

		w.Write(data)
	})

	mx.Path("/frame/{timestamp}").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts, err := parseTimestamp(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
			return
		}

		log.Printf("search video bundle with timestamp: %s", ts.Format(time.RFC1123))

		bundle, err := bundler.SearchVideoBundle(ts)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		frame, err := encoder.ExtractFrame(bundle, ts.Second())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "video encoding failed: %s", err)
			return
		}

		w.Header().Add("Content-Type", "image/jpeg")
		w.Header().Add("Content-Length", fmt.Sprint(len(frame)))

		w.WriteHeader(http.StatusOK)

		w.Write(frame)
	})
}

func parseTimestamp(r *http.Request) (ts time.Time, err error) {
	vars := mux.Vars(r)
	timestamp := vars["timestamp"]

	if timestamp == "" {
		err = errors.New("no timestamp")
		return
	}

	ts, err = time.ParseInLocation("20060102150405", timestamp, time.Local)
	if err != nil {
		err = fmt.Errorf("timestamp parse failed: %w", err)
		return
	}

	return
}

package server

import (
	"fmt"
	"html/template"
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

		tpl := template.Must(template.New("index").Parse(`
			<!DOCTYPE html>
			<html>
				<head>
					<title>Event video</title>
					<style>
						body {
							margin: 0;
							overflow: hidden;
						}
						video {
							width: 100vw;
							height: 100vh;
							padding: 1em;
							box-sizing: border-box;
						}
					</style>
					<script>
						window.addEventListener('load', () => {
							const video = document.getElementById("video");
							video.addEventListener('click', () => video.play());
						});
					</script>
				</head>
				<body>
					<video id="video" src="/event/{{.}}" />
				</body>
			</html>
		`))

		err := tpl.Execute(w, timestamp)
		if err != nil {
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

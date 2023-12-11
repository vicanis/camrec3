package stream

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func StartHealthcheck(ctx context.Context) chan error {
	ch := make(chan error)

	go func() {
		srv := http.Server{
			Addr:         "127.0.0.1:8080",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if err := checkProcessState(); err != nil {
					http.Error(w, "bad process state", http.StatusInternalServerError)
				} else {
					fmt.Fprint(w, "ok")
				}
			}),
		}

		select {
		case <-ctx.Done():
			ch <- srv.Shutdown(ctx)
		case ch <- srv.ListenAndServe():
		}
	}()

	return ch
}

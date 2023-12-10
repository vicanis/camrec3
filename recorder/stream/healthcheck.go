package stream

import (
	"context"
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
				cmdLock.RLock()
				defer cmdLock.RUnlock()

				if err := checkProcessState(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}),
		}

		select {
		case <-ctx.Done():
			srv.Shutdown(ctx)
		case ch <- srv.ListenAndServe():
		}
	}()

	return ch
}

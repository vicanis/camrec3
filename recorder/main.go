package main

import (
	"camrec/stream"
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if os.Getenv("STREAM") == "" {
		log.Fatal("no stream URL")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	go func() {
		ech := stream.Start(ctx)

		select {
		case <-ctx.Done():
			return

		case err := <-ech:
			log.Printf("streaming end: %s", err)
			cancel()
			return
		}
	}()

	go func() {
		ech := stream.StartHealthcheck(ctx)

		select {
		case <-ctx.Done():
		case err := <-ech:
			log.Printf("healtcheck failed: %s", err)
			cancel()
		}
	}()

	time.Sleep(time.Second)

	log.Printf("press ctrl+c to interrupt")

	select {
	case sig := <-sigchan:
		log.Printf("signal: %s", sig)
		cancel()
	case <-ctx.Done():
		log.Printf("terminated: %s", ctx.Err())
	}

	time.Sleep(time.Second)
}

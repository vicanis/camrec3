package main

import (
	"log"
	"processor/server"
)

func main() {
	log.Fatal(server.Start())
}

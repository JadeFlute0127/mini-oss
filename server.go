package main

import (
	"log"
	"net/http"
	"os"
	"version_1/objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}

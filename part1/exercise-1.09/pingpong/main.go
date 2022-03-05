package main

import (
	"log"
	"net/http"
)

func main() {
	server := Server{}
	log.Fatal(http.ListenAndServe(":4000", &server))
}
